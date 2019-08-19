package redisai

import (
	"fmt"
	"io/ioutil"
	"reflect"
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

type PipelinedClient struct {
	Pool            *redis.Pool
	PipelineMaxSize int
	PipelinePos     int
	ActiveConn      redis.Conn
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


// ConnectPipelined intializes a Client with pipeline enabled by default
func ConnectPipelined(url string, pipelineMax int ) (c *PipelinedClient) {
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
	defer func() { if c.ActiveConn != nil { c.ActiveConn.Flush(); c.ActiveConn.Close() } }()
	return c
}

// Close ensures that no connection is kept alive and prior to that we flush all db commands
func (c *PipelinedClient) Close()  {
	if c.ActiveConn != nil { c.ActiveConn.Flush(); c.ActiveConn.Close() }
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
	args := TensorSetArgs(name, dt, dims, data, false )
	if args == nil {
		return fmt.Errorf("redisai.TensorSet: unknown type %T",reflect.TypeOf(data))
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

// TensorSet sets a tensor
func (c *PipelinedClient) TensorSet(name string, dt DataType, dims []int, data interface{}) (err error) {
	args := TensorSetArgs(name, dt, dims, data, false )
	if args == nil {
		return fmt.Errorf("redisai.TensorSet: unknown type %T",reflect.TypeOf(data))
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

func (c *PipelinedClient) forceFlush()  (err error) {
	err = nil
	if c.ActiveConn != nil {
		c.PipelinePos = 0
		err = c.ActiveConn.Flush()
	}
return err
}

func (c *PipelinedClient) pipeIncr(conn redis.Conn)  (err error) {
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

// TensorGetValues gets a tensor's values
func (c *Client) TensorGetValues(name string) (dt DataType, shape []int, data []float64, err error) {
	args := redis.Args{}.Add(name, "VALUES")
	conn := c.pool.Get()
	defer conn.Close()

	rep, err := redis.MultiBulk(conn.Do("AI.TENSORGET", args...))
	if err != nil {
		return
	}
	return ProcessTensorResponse(rep)
}

func TensorSetArgs(name string, dt DataType, dims []int, data interface{}, includeCommandName bool ) redis.Args {
	args := redis.Args{}
	if includeCommandName {
		args = args.Add( "AI.TENSORSET" )
	}
	args = args.Add(name, dt).AddFlat(dims)
	var dtype = reflect.TypeOf(data)
	switch dtype {
	case reflect.TypeOf(([]byte)(nil)):
		args = args.Add("BLOB", data)
	case reflect.TypeOf((string)("")):
		fallthrough
	case reflect.TypeOf(([]int)(nil)):
		fallthrough
	case reflect.TypeOf(([]int8)(nil)):
		fallthrough
	case reflect.TypeOf(([]int16)(nil)):
		fallthrough
	case reflect.TypeOf(([]int32)(nil)):
		fallthrough
	case reflect.TypeOf(([]int64)(nil)):
		fallthrough
	case reflect.TypeOf(([]uint)(nil)):
		fallthrough
	case reflect.TypeOf(([]uint16)(nil)):
		fallthrough
	case reflect.TypeOf(([]uint32)(nil)):
		fallthrough
	case reflect.TypeOf(([]uint64)(nil)):
		fallthrough
	case reflect.TypeOf(([]float32)(nil)):
		fallthrough
	case reflect.TypeOf(([]float64)(nil)):
		args = args.Add("VALUES").AddFlat(data)
	default:
		args = nil
	}
	return args
}


func ModelRunArgs(name string, inputs []string, outputs []string, includeCommandName bool) redis.Args {
	args := redis.Args{}
	if includeCommandName {
		args = args.Add( "AI.MODELRUN" )
	}
	args = args.Add(name)
	if len(inputs) > 0 {
		args = args.Add("INPUTS").AddFlat(inputs)
	}
	if len(outputs) > 0 {
		args = args.Add("OUTPUTS").AddFlat(outputs)
	}
	return args
}


func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i:=0; i<s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func ProcessTensorResponse(respInitial interface{})( dt DataType, shape []int, data []float64, err error ) {
	rep := InterfaceSlice(respInitial)

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
