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
		wantErr bool
	}{
		{"test:TestTensorSetArgs:[]float32:1", args{"test:TestTensorSetArgs:1", TypeFloat, []int{1}, []float32{1}, true}, string(TensorContentTypeValues),false},
		{"test:TestTensorSetArgs:[]byte:1", args{"test:TestTensorSetArgs:1", TypeFloat, []int{1}, f32Bytes, true}, string(TensorContentTypeBlob),false},
		{"test:TestTensorSetArgs:[]int:1", args{"test:TestTensorSetArgs:1", TypeInt32, []int{1}, []int{1}, true}, string(TensorContentTypeValues),false},
		{"test:TestTensorSetArgs:[]int8:1", args{"test:TestTensorSetArgs:1", TypeInt8, []int{1}, []int8{1}, true}, string(TensorContentTypeValues),false},
		{"test:TestTensorSetArgs:[]int16:1", args{"test:TestTensorSetArgs:1", TypeInt16, []int{1}, []int16{1}, true}, string(TensorContentTypeValues),false},
		{"test:TestTensorSetArgs:[]int64:1", args{"test:TestTensorSetArgs:1", TypeInt64, []int{1}, []int64{1}, true}, string(TensorContentTypeValues),false},
		{"test:TestTensorSetArgs:[]uint8:1", args{"test:TestTensorSetArgs:1", TypeUint8, []int{1}, []uint8{1}, true}, string(TensorContentTypeBlob),false},
		{"test:TestTensorSetArgs:[]uint16:1", args{"test:TestTensorSetArgs:1", TypeUint16, []int{1}, []uint16{1}, true}, string(TensorContentTypeValues),false},
		{"test:TestTensorSetArgs:[]uint32:1", args{"test:TestTensorSetArgs:1", TypeUint8, []int{1}, []uint32{1}, true}, string(TensorContentTypeBlob),true},
		{"test:TestTensorSetArgs:[]uint64:1", args{"test:TestTensorSetArgs:1", TypeUint16, []int{1}, []uint64{1}, true}, string(TensorContentTypeValues),true},
		{"test:TestTensorSetArgs:[]float32:1", args{"test:TestTensorSetArgs:1", TypeFloat32, []int{1}, []float32{1}, true}, string(TensorContentTypeValues),false},
		{"test:TestTensorSetArgs:[]float64:1", args{"test:TestTensorSetArgs:1", TypeFloat64, []int{1}, []float64{1}, true}, string(TensorContentTypeValues),false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got,err := TensorSetArgs(tt.args.name, tt.args.dt, tt.args.dims, tt.args.data, tt.args.includeCommandName)

			if (err != nil) != tt.wantErr {
				t.Errorf("TensorSetArgs() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr == false {
				if got[4].(string) != tt.want {
					t.Errorf("TensorSetArgs() TensorContentType = %v, want %v", got[4], tt.want)
				}
			}
		})
	}
}
