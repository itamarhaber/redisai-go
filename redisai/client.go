package redisai

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Client is a RedisAI client
type Client struct {
	pool *redis.Pool
}

func (c *Client) TensorGet(name string, ct TensorContentType) (data []interface{}, err error) {
	args := redis.Args{}.Add(name, ct)
	conn := c.pool.Get()
	defer conn.Close()
	resp, err := conn.Do("AI.TENSORGET", args...)
	data, err = processTensorReplyMeta(resp,err)
	if ct == TensorContentTypeBlob {
		data, err = processTensorReplyBlob(data,err)
	}
	if ct == TensorContentTypeValues {
		data, err = processTensorReplyValues(data,err)
	}
	return
}


// TensorGetValues gets a tensor's values
func (c *Client) TensorGetValues(name string) (dt DataType, shape []int, data []interface{}, err error) {
	args := redis.Args{}.Add(name, TensorContentTypeValues)
	conn := c.pool.Get()
	defer conn.Close()

	rep, err := conn.Do("AI.TENSORGET", args...)
	if err != nil {
		return
	}
	resp,err := processTensorReplyMeta(rep,err)
	if err != nil {
		return
	}
	resp,err = processTensorReplyValues(resp,err)
	if err != nil {
		return
	}
	return resp[0].(DataType), resp[1].([]int), resp[2].([]interface{}), err
}

// TensorGetValues gets a tensor's values
func (c *Client) TensorGetMeta(name string) (dt DataType, shape []int, err error) {
	args := redis.Args{}.Add(name, TensorContentTypeMeta)
	conn := c.pool.Get()
	defer conn.Close()
	rep, err := conn.Do("AI.TENSORGET", args...)
	if err != nil {
		return
	}
	resp,err := processTensorReplyMeta(rep,err)
	if err != nil {
		return
	}
	return resp[0].(DataType), resp[1].([]int), err
}

// TensorGetValues gets a tensor's values
func (c *Client) TensorGetBlob(name string) (dt DataType, shape []int, data []byte, err error) {
	args := redis.Args{}.Add(name, TensorContentTypeBlob)
	conn := c.pool.Get()
	defer conn.Close()

	rep, err := conn.Do("AI.TENSORGET", args...)
	if err != nil {
		return
	}

	resp,err := processTensorReplyMeta(rep,err)
	if err != nil {
		return
	}
	resp,err = processTensorReplyBlob(resp,err)
	if err != nil {
		return
	}

	return resp[0].(DataType), resp[1].([]int), resp[2].([]byte), err
}

func (c *Client) ModelGet(name string) (data []byte, err error) {
	args := redis.Args{}.Add(name)
	conn := c.pool.Get()
	defer conn.Close()
	respInitial, err := conn.Do("AI.MODELGET", args...)
	if err != nil {
		return
	}
	rep, err := redis.Values(respInitial, err)
	if len(rep) != 3 {
		err = fmt.Errorf("redisai.ModelGet: AI.MODELGET returned response with incorrect sizing. expected '%d' got '%d'", 3, len(rep))
		return
	}
	data, err = redis.Bytes(rep[2], nil)
	if err != nil {
		return
	}
	return
}

func (c *Client) ModelDel(name string) (err error) {
	args := redis.Args{}.Add(name)
	conn := c.pool.Get()
	defer conn.Close()
	_, err = conn.Do("AI.MODELDEL", args...)
	return
}

func (c *Client) ScriptGet(name string) (data []byte, err error) {
	args := redis.Args{}.Add(name)
	conn := c.pool.Get()
	defer conn.Close()
	respInitial, err := conn.Do("AI.SCRIPTGET", args...)
	if err != nil {
		return
	}
	rep, err := redis.Values(respInitial, err)
	if len(rep) != 3 {
		err = fmt.Errorf("redisai.ScriptGet: AI.SCRIPTGET returned response with incorrect sizing. expected '%d' got '%d'", 3, len(rep))
		return
	}
	data, err = redis.Bytes(rep[2], nil)
	if err != nil {
		return
	}
	return
}

func (c *Client) ScriptDel(name string) (err error) {
	args := redis.Args{}.Add(name)
	conn := c.pool.Get()
	defer conn.Close()
	rep, err := redis.String(conn.Do("AI.SCRIPTDEL", args...))
	if err != nil {
		return err
	}
	if rep != "OK" {
		return fmt.Errorf("redisai.ScriptDel: AI.SCRIPTDEL returned '%s'", rep)
	}
	return
}

func (c *Client) LoadBackend(backend_identifier BackendType, location string) (err error) {
	args := redis.Args{}.Add("LOADBACKEND").Add(backend_identifier).Add(location)
	conn := c.pool.Get()
	defer conn.Close()
	rep, err := redis.String(conn.Do("AI.CONFIG", args...))
	if err != nil {
		return err
	}
	if rep != "OK" {
		return fmt.Errorf("redisai.LoadBackend: AI.CONFIG LOADBACKEND returned '%s'", rep)
	}
	return
}

// Connect intializes a Client
func Connect(url string, pool *redis.Pool) (c *Client) {
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

	c = &Client{
		pool: cpool,
	}
	return c
}

// ModelSet sets a RedisAI model from a blob
func (c *Client) ModelSet(name string, backend BackendType, device DeviceType, data []byte, inputs []string, outputs []string) error {
	args := redis.Args{}.Add(name, backend, device)
	if len(inputs) > 0 {
		args = args.Add("INPUTS").AddFlat(inputs)
	}
	if len(outputs) > 0 {
		args = args.Add("OUTPUTS").AddFlat(outputs)
	}
	args = args.Add(data)

	conn := c.pool.Get()
	defer conn.Close()
	rep, err := redis.String(conn.Do("AI.MODELSET", args...))
	if err != nil {
		return err
	}
	if rep != "OK" {
		return fmt.Errorf("redisai.ModelSet: AI.MODELSET returned '%s'", rep)
	}
	return nil
}

// ModelSetFromFile sets a RedisAI model from a file
func (c *Client) ModelSetFromFile(name string, backend BackendType, device DeviceType, path string, inputs []string, outputs []string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return c.ModelSet(name, backend, device, data, inputs, outputs)
}

// ModelRun runs a RedisAI model
func (c *Client) ModelRun(name string, inputs []string, outputs []string) error {
	args := ModelRunArgs(name, inputs, outputs, false)
	conn := c.pool.Get()
	defer conn.Close()

	rep, err := redis.String(conn.Do("AI.MODELRUN", args...))
	if err != nil {
		return err
	}
	if rep != "OK" {
		return fmt.Errorf("redisai.ModelRun: AI.MODELRUN returned '%s'", rep)
	}
	return nil
}

// ScriptSet sets a RedisAI script from a blob
func (c *Client) ScriptSet(name string, device DeviceType, data []byte) error {
	args := redis.Args{}.Add(name, device, data)

	conn := c.pool.Get()
	defer conn.Close()
	rep, err := redis.String(conn.Do("AI.SCRIPTSET", args...))
	if err != nil {
		return err
	}
	if rep != "OK" {
		return fmt.Errorf("redisai.ScriptSet: AI.SCRIPTSET returned '%s'", rep)
	}
	return nil
}

// ScriptSetFromFile sets a RedisAI script from a file
func (c *Client) ScriptSetFromFile(name string, device DeviceType, path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return c.ScriptSet(name, device, data)
}

// ScriptRun runs a RedisAI script
func (c *Client) ScriptRun(name string, fn string, inputs []string, outputs []string) error {
	args := redis.Args{}.Add(name, fn)
	if len(inputs) > 0 {
		args = args.Add("INPUTS").AddFlat(inputs)
	}
	if len(outputs) > 0 {
		args = args.Add("OUTPUTS").AddFlat(outputs)
	}
	conn := c.pool.Get()
	defer conn.Close()

	rep, err := redis.String(conn.Do("AI.SCRIPTRUN", args...))
	if err != nil {
		return err
	}
	if rep != "OK" {
		return fmt.Errorf("redisai.ScriptRun: AI.SCRIPTRUN returned '%s'", rep)
	}
	return nil
}

// TensorSet sets a tensor
func (c *Client) TensorSet(name string, dt DataType, dims []int, data interface{}) (err error) {
	args := TensorSetArgs(name, dt, dims, data, false)
	if args == nil {
		return fmt.Errorf("redisai.TensorSet: unknown type %T", reflect.TypeOf(data))
	}
	conn := c.pool.Get()
	defer conn.Close()

	rep, err := redis.String(conn.Do("AI.TENSORSET", args...))
	if err != nil {
		return err
	}
	if rep != "OK" {
		return fmt.Errorf("redisai.TensorSet: AI.TENSORSET returned '%s'", rep)
	}
	return nil
}
