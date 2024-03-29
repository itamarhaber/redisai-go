package redisai

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gomodule/redigo/redis"
)

// DeviceType is a device type
type DeviceType string

// BackendType is a backend type
type BackendType string

// DataType is a data type
type DataType string

const (
	// BackendTF represents a TensorFlow backend
	BackendTF = BackendType("TF")
	// BackendTorch represents a Torch backend
	BackendTorch = BackendType("TORCH")
	// BackendONNX represents an ONNX backend
	BackendONNX = BackendType("ORT")

	// DeviceCPU represents a CPU device
	DeviceCPU = DeviceType("CPU")
	// DeviceGPU represents a GPU device
	DeviceGPU = DeviceType("GPU")

	// TypeFloat represents a float type
	TypeFloat = DataType("FLOAT")
	// TypeDouble represents a double type
	TypeDouble = DataType("DOUBLE")
	// TypeInt8 represents a int8 type
	TypeInt8 = DataType("INT8")
	// TypeInt16 represents a int16 type
	TypeInt16 = DataType("INT16")
	// TypeInt32 represents a int32 type
	TypeInt32 = DataType("INT32")
	// TypeInt64 represents a int64 type
	TypeInt64 = DataType("INT64")
	// TypeUint8 represents a uint8 type
	TypeUint8 = DataType("UINT8")
	// TypeUint16 represents a uint16 type
	TypeUint16 = DataType("UINT16")
	// TypeUint32 represents a uint32 type
	TypeUint32 = DataType("UINT32")
	// TypeUint64 represents a uint64 type
	TypeUint64 = DataType("UINT64")
	// TypeFloat32 is an alias for float
	TypeFloat32 = DataType("FLOAT")
	// TypeFloat64 is an alias for double
	TypeFloat64 = DataType("DOUBLE")
)

// Client is a RedisAI client
type Client struct {
	pool *redis.Pool
}

// Connect intializes a Client
func Connect(url string) (c *Client) {
	c = &Client{
		pool: &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial:        func() (redis.Conn, error) { return redis.DialURL(url) },
		},
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
	args := redis.Args{}.Add(name)
	if len(inputs) > 0 {
		args = args.Add("INPUTS").AddFlat(inputs)
	}
	if len(outputs) > 0 {
		args = args.Add("OUTPUTS").AddFlat(outputs)
	}
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
	args := redis.Args{}.Add(name, dt).AddFlat(dims)
	switch v := data.(type) {
	case string:
	case []byte:
		args = args.Add("BLOB", v)
	case []int:
	case []int8:
	case []int16:
	case []int32:
	case []int64:
	case []uint:
	case []uint16:
	case []uint32:
	case []uint64:
	case []float32:
	case []float64:
		args = args.Add("VALUES").AddFlat(v)
	default:
		return fmt.Errorf("redisai.TensorSet: unknown type %T", v)
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

// TensorGetValues gets a tensor's values
func (c *Client) TensorGetValues(name string) (dt DataType, shape []int, data []float64, err error) {
	args := redis.Args{}.Add(name, "VALUES")
	conn := c.pool.Get()
	defer conn.Close()

	rep, err := redis.MultiBulk(conn.Do("AI.TENSORGET", args...))
	if err != nil {
		return
	}

	sdt, err := redis.String(rep[0], nil)
	if err != nil {
		return
	}
	shape, err = redis.Ints(rep[1], nil)
	if err != nil {
		return
	}
	data, err = redis.Float64s(rep[2], nil)
	if err != nil {
		return
	}

	switch sdt {
	case "FLOAT":
		dt = TypeFloat
	case "DOUBLE":
		dt = TypeDouble
	case "INT8":
		dt = TypeInt8
	case "INT16":
		dt = TypeInt16
	case "INT32":
		dt = TypeInt32
	case "INT64":
		dt = TypeInt64
	case "UINT8":
		dt = TypeUint8
	case "UINT16":
		dt = TypeUint16
	case "UINT32":
		dt = TypeUint32
	case "UINT64":
		dt = TypeUint64
	default:
		err = fmt.Errorf("redisai.TensorGet: AI.TENSORGET returned unknown type '%s'", sdt)
		return
	}
	return
}
