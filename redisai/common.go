package redisai

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"reflect"
)

type aiclient interface {
	area() float64
	perim() float64
}

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

func TensorSetArgs(name string, dt DataType, dims []int, data interface{}, includeCommandName bool) redis.Args {
	args := redis.Args{}
	if includeCommandName {
		args = args.Add("AI.TENSORSET")
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
		args = args.Add("AI.MODELRUN")
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

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func ProcessTensorResponse(respInitial interface{}) (dt DataType, shape []int, data []float64, err error) {
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
