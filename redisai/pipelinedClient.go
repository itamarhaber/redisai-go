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

func (c *PipelinedClient) ModelSet(name string, backend BackendType, device DeviceType, data []byte, inputs []string, outputs []string) ( err error) {
	args := redis.Args{}.Add(name, backend, device)
	if len(inputs) > 0 {
		args = args.Add("INPUTS").AddFlat(inputs)
	}
	if len(outputs) > 0 {
		args = args.Add("OUTPUTS").AddFlat(outputs)
	}
	args = args.Add(data)

	if c.ActiveConn == nil {
		c.ActiveConn = c.Pool.Get()
		defer c.ActiveConn.Close()
	}
	err = c.ActiveConn.Send("AI.MODELSET", args...)
	if err != nil {
		return err
	}
	// incremement the pipeline
	// flush if required
	err = c.pipeIncr(c.ActiveConn)
	if err != nil {
		return err
	}
	return nil
}

func (c *PipelinedClient) ScriptSet(name string, device DeviceType, script_source string) (err error) {
	args := redis.Args{}.Add(name, device, script_source)
	err = c.SendAndIncr("AI.SCRIPTSET" , args)
	if err != nil {
		return err
	}
	return nil
}

func (c *PipelinedClient) SendAndIncr( commandName string, args redis.Args) (err error) {
	if c.ActiveConn == nil {
		c.ActiveConn = c.Pool.Get()
	}
	err = c.ActiveConn.Send(commandName, args...)
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

func (c *PipelinedClient) ScriptRun(name string, fn string, inputs []string, outputs []string) (err error) {
	args := redis.Args{}.Add(name, fn)
	if len(inputs) > 0 {
		args = args.Add("INPUTS").AddFlat(inputs)
	}
	if len(outputs) > 0 {
		args = args.Add("OUTPUTS").AddFlat(outputs)
	}
	err = c.SendAndIncr("AI.SCRIPTRUN" , args)
	if err != nil {
		return err
	}
	return
}

func (c *PipelinedClient) ModelGet(name string) (data []interface{}, err error) {
	args := redis.Args{}.Add(name)
	err = c.SendAndIncr("AI.MODELGET" , args)
	if err != nil {
		return data, err
	}
	return
}

func (c *PipelinedClient) ModelDel(name string) (err error) {
	args := redis.Args{}.Add(name)
	err = c.SendAndIncr("AI.MODELDEL" , args)
	if err != nil {
		return err
	}
	return
}

func (c *PipelinedClient) ScriptGet(name string) (data []interface{}, err error) {
	args := redis.Args{}.Add(name)
	err = c.SendAndIncr("AI.SCRIPTGET" , args)
	if err != nil {
		return data, err
	}
	return
}

func (c *PipelinedClient) ScriptDel(name string) (err error) {
	args := redis.Args{}.Add(name)
	err = c.SendAndIncr("AI.SCRIPTDEL" , args)
	if err != nil {
		return err
	}
	return
}

func (c *PipelinedClient) LoadBackend(backend_identifier BackendType, location string) (err error) {
	args := redis.Args{}.Add("LOADBACKEND").Add(backend_identifier).Add(location)
	err = c.SendAndIncr("AI.CONFIG" , args)
	if err != nil {
		return err
	}
	return
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
	err = c.SendAndIncr("AI.MODELRUN" , args)
	if err != nil {
		return err
	}
	return
}

// TensorSet sets a tensor
func (c *PipelinedClient) TensorSet(name string, dt DataType, shape []int, data interface{}) (err error) {
	args,err := TensorSetArgs(name, dt, shape, data, false)
	err = c.SendAndIncr("AI.TENSORSET" , args)
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