package redisai

import (
	"fmt"
	"reflect"
	"time"

	"github.com/gomodule/redigo/redis"
)

type PipelinedClient struct {
	Pool            *redis.Pool
	PipelineMaxSize int
	PipelinePos     int
	ActiveConn      redis.Conn
}

// ConnectPipelined intializes a Client with pipeline enabled by default
func ConnectPipelined(url string, pipelineMax int) (c *PipelinedClient) {
	c = &PipelinedClient{
		Pool: &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial:        func() (redis.Conn, error) { return redis.DialURL(url) },
		},
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
func (c *PipelinedClient) Close() {
	if c.ActiveConn != nil {
		c.ActiveConn.Flush()
		c.ActiveConn.Close()
	}
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
	return nil
}

// TensorSet sets a tensor
func (c *PipelinedClient) TensorSet(name string, dt DataType, dims []int, data interface{}) (err error) {
	args := TensorSetArgs(name, dt, dims, data, false)
	if args == nil {
		return fmt.Errorf("redisai.TensorSet: unknown type %T", reflect.TypeOf(data))
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
	return nil
}

func (c *PipelinedClient) forceFlush() (err error) {
	err = nil
	if c.ActiveConn != nil {
		c.PipelinePos = 0
		err = c.ActiveConn.Flush()
	}
	return err
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
	return nil
}

// TensorGetValues gets a tensor's values
func (c *PipelinedClient) TensorGetValues(name string) (err error) {
	args := redis.Args{}.Add(name, "VALUES")

	if c.ActiveConn == nil {
		c.ActiveConn = c.Pool.Get()
		defer c.ActiveConn.Close()
	}
	err = c.ActiveConn.Send("AI.TENSORGET", args...)
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
