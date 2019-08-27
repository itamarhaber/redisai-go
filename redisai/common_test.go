package redisai

import (
	"github.com/gomodule/redigo/redis"
	"reflect"
	"testing"
)

func TestModelRunArgs(t *testing.T) {
	type args struct {
		name               string
		inputs             []string
		outputs            []string
		includeCommandName bool
	}
	tests := []struct {
		name string
		args args
		want redis.Args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ModelRunArgs(tt.args.name, tt.args.inputs, tt.args.outputs, tt.args.includeCommandName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ModelRunArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseTensorResponseBlob(t *testing.T) {
	type args struct {
		respInitial interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantDt    DataType
		wantShape []int
		wantData  []byte
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDt, gotShape, gotData, err := ParseTensorResponseBlob(tt.args.respInitial)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTensorResponseBlob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDt != tt.wantDt {
				t.Errorf("ParseTensorResponseBlob() gotDt = %v, want %v", gotDt, tt.wantDt)
			}
			if !reflect.DeepEqual(gotShape, tt.wantShape) {
				t.Errorf("ParseTensorResponseBlob() gotShape = %v, want %v", gotShape, tt.wantShape)
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("ParseTensorResponseBlob() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestParseTensorResponseMeta(t *testing.T) {
	type args struct {
		respInitial interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantDt    DataType
		wantShape []int
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDt, gotShape, err := ParseTensorResponseMeta(tt.args.respInitial)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTensorResponseMeta() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDt != tt.wantDt {
				t.Errorf("ParseTensorResponseMeta() gotDt = %v, want %v", gotDt, tt.wantDt)
			}
			if !reflect.DeepEqual(gotShape, tt.wantShape) {
				t.Errorf("ParseTensorResponseMeta() gotShape = %v, want %v", gotShape, tt.wantShape)
			}
		})
	}
}

func TestParseTensorResponseValues(t *testing.T) {
	type args struct {
		respInitial interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantDt    DataType
		wantShape []int
		wantData  interface{}
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDt, gotShape, gotData, err := ParseTensorResponseValues(tt.args.respInitial)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTensorResponseValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDt != tt.wantDt {
				t.Errorf("ParseTensorResponseValues() gotDt = %v, want %v", gotDt, tt.wantDt)
			}
			if !reflect.DeepEqual(gotShape, tt.wantShape) {
				t.Errorf("ParseTensorResponseValues() gotShape = %v, want %v", gotShape, tt.wantShape)
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("ParseTensorResponseValues() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestTensorSetArgs_TensorContentType(t *testing.T) {

	f32Bytes, _ := float32ToByte(1.1)

	type args struct {
		name               string
		dt                 DataType
		dims               []int
		data               interface{}
		includeCommandName bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test:TestTensorSetArgs:[]float32:1", args{"test:TestTensorSetArgs:1", TypeFloat, []int{1}, []float32{1}, true}, string(TensorContentTypeValues)},
		{"test:TestTensorSetArgs:[]byte:1", args{"test:TestTensorSetArgs:1", TypeFloat, []int{1}, f32Bytes, true}, string(TensorContentTypeBlob)},
		{"test:TestTensorSetArgs:[]int:1", args{"test:TestTensorSetArgs:1", TypeInt32, []int{1}, []int{1}, true}, string(TensorContentTypeValues)},
		{"test:TestTensorSetArgs:[]int8:1", args{"test:TestTensorSetArgs:1", TypeInt8, []int{1}, []int8{1}, true}, string(TensorContentTypeValues)},
		{"test:TestTensorSetArgs:[]int16:1", args{"test:TestTensorSetArgs:1", TypeInt16, []int{1}, []int16{1}, true}, string(TensorContentTypeValues)},
		{"test:TestTensorSetArgs:[]int64:1", args{"test:TestTensorSetArgs:1", TypeInt64, []int{1}, []int64{1}, true}, string(TensorContentTypeValues)},
		{"test:TestTensorSetArgs:[]uint:1", args{"test:TestTensorSetArgs:1", TypeUint32, []int{1}, []uint32{1}, true}, string(TensorContentTypeValues)},
		{"test:TestTensorSetArgs:[]uint8:1", args{"test:TestTensorSetArgs:1", TypeUint8, []int{1}, []uint8{1}, true}, string(TensorContentTypeBlob)},
		{"test:TestTensorSetArgs:[]uint16:1", args{"test:TestTensorSetArgs:1", TypeUint16, []int{1}, []uint16{1}, true}, string(TensorContentTypeValues)},
		{"test:TestTensorSetArgs:[]uint32:1", args{"test:TestTensorSetArgs:1", TypeUint32, []int{1}, []uint32{1}, true}, string(TensorContentTypeValues)},
		{"test:TestTensorSetArgs:[]uint64:1", args{"test:TestTensorSetArgs:1", TypeUint64, []int{1}, []uint64{1}, true}, string(TensorContentTypeValues)},
		{"test:TestTensorSetArgs:[]float32:1", args{"test:TestTensorSetArgs:1", TypeFloat32, []int{1}, []float32{1}, true}, string(TensorContentTypeValues)},
		{"test:TestTensorSetArgs:[]float64:1", args{"test:TestTensorSetArgs:1", TypeFloat64, []int{1}, []float64{1}, true}, string(TensorContentTypeValues)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TensorSetArgs(tt.args.name, tt.args.dt, tt.args.dims, tt.args.data, tt.args.includeCommandName)
			if got[4].(string) != tt.want {
				t.Errorf("TensorSetArgs() TensorContentType = %v, want %v", got[4], tt.want)
			}
		})
	}
}
