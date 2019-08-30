package redisai

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type PipelinedClient struct {
	Pool            *redis.Pool
	PipelineMaxSize int
	PipelinePos     int
	ActiveConn      redis.Conn
}

func (c *PipelinedClient) TensorGet(name string, ct TensorContentType) (data []interface{}, err error)  {
	data = nil
	args := redis.Args{}.Add(name, ct)

	if c.ActiveConn == nil {
		c.ActiveConn = c.Pool.Get()
		defer c.ActiveConn.Close()
	}
	err = c.ActiveConn.Send("AI.TENSORGET", args...)
	if err != nil {
		return data, err
	}
	// incremement the pipeline
	// flush if required
	err = c.pipeIncr(c.ActiveConn)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (c *PipelinedClient) ModelSet(name string, backend BackendType, device DeviceType, data []byte, inputs []string, outputs []string) error {
	panic("implement me")
}

func (c *PipelinedClient) ScriptSet(name string, device DeviceType, data string) error {
	panic("implement me")
}

func (c *PipelinedClient) ScriptRun(name string, fn string, inputs []string, outputs []string) error {
	panic("implement me")
}

func (c *PipelinedClient) ModelGet(name string) (data []interface{}, err error) {
	panic("implement me")
}

func (c *PipelinedClient) ModelDel(name string) (err error) {
	panic("implement me")
}

func (c *PipelinedClient) ScriptGet(name string) (data []interface{}, err error) {
	panic("implement me")
}

func (c *PipelinedClient) ScriptDel(name string) (err error) {
	panic("implement me")
}

func (c *PipelinedClient) LoadBackend(backendIdentifier BackendType, location string) (err error) {
	panic("implement me")
}

// ConnectPipelined intializes a Client with pipeline enabled by default
func ConnectPipelined(url string, pipelineMax int, pool *redis.Pool) (c *PipelinedClient) {

	var cpool *redis.Pool = nil
	if pool == nil {
		cpool = &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial:        func() (redis.Conn, error) { return redis.DialURL(url) },
		}
	} else {
		cpool = pool
	}

	c = &PipelinedClient{
		Pool:            cpool,
		PipelineMaxSize: pipelineMax,
		PipelinePos:     0,
		ActiveConn:      nil,
	}
	defer func() {
		if c.ActiveConn != nil {
			c.ActiveConn.Flush()
			c.ActiveConn.Close()
		}
	}()
	return c
}

// Close ensures that no connection is kept alive and prior to that we flush all db commands
func (c *PipelinedClient) Close() (err error) {
	if c.ActiveConn != nil {
		err = c.ActiveConn.Flush()
		if err != nil {
			return
		}
		err = c.ActiveConn.Close()
		if err != nil {
			return
		}
	}

	return
}

// ModelRun runs a RedisAI model
func (c *PipelinedClient) ModelRun(name string, inputs []string, outputs []string) (err error) {
	args := ModelRunArgs(name, inputs, outputs, false)
	if c.ActiveConn == nil {
		c.ActiveConn = c.Pool.Get()
		defer c.ActiveConn.Close()
	}
	err = c.ActiveConn.Send("AI.MODELRUN", args...)
	if err != nil {
		return err
	}
	// incremement the pipeline
	// flush if required
	err = c.pipeIncr(c.ActiveConn)
	if err != nil {
		return err
	}
	return
}

// TensorSet sets a tensor
func (c *PipelinedClient) TensorSet(name string, dt DataType, shape []int, data interface{}) (err error) {
	args,err := TensorSetArgs(name, dt, shape, data, false)
	if err != nil {
		return err
	}
	if c.ActiveConn == nil {
		c.ActiveConn = c.Pool.Get()
	}
	err = c.ActiveConn.Send("AI.TENSORSET", args...)
	if err != nil {
		return err
	}
	// incremement the pipeline
	// flush if required
	err = c.pipeIncr(c.ActiveConn)
	if err != nil {
		return err
	}
	return
}

func (c *PipelinedClient) ForceFlush() (err error) {
	err = nil
	if c.ActiveConn != nil {
		c.PipelinePos = 0
		err = c.ActiveConn.Flush()
	}
	return
}

func (c *PipelinedClient) pipeIncr(conn redis.Conn) (err error) {
	c.PipelinePos++
	if c.PipelinePos >= c.PipelineMaxSize {
		err = conn.Flush()
		c.PipelinePos = 0
	}
	if err != nil {
		return err
	}
	return
}