package redisai

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"reflect"
	"testing"
)

func TestModelRunArgs(t *testing.T) {
	nameT1 := "test:ModelRunArgs:1:includeCommandName"
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
		{ nameT1, args{ nameT1, []string{}, []string{}, true }, redis.Args{ "AI.MODELRUN", nameT1 } },
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

func Test_replyDataType(t *testing.T) {

	var r1 interface{} = string("abc")
	var r2 interface{} = int(1)
	var r3 interface{} = string("FLOAT")
	var r4 interface{} = string("DOUBLE")
	var r5 interface{} = string("INT8")
	var r6 interface{} = string("INT16")
	var r7 interface{} = string("INT32")
	var r8 interface{} = string("INT64")
	var r9 interface{} = string("UINT8")
	var r10 interface{} = string("UINT16")

	var err1 error = fmt.Errorf("")

	type args struct {
		reply   interface{}
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantDt  DataType
		wantErr bool
	}{
		{   "test:replyDataType:Error:1",  args{ r1, err1 } , "", true },
		{   "test:replyDataType:Error:WrongType:2",  args{ r2, nil } , "", true },
		{   "test:replyDataType:FLOAT:3",  args{ r3, nil } , TypeFloat, false },
		{   "test:replyDataType:DOUBLE:4",  args{ r4, nil } , TypeDouble, false },
		{   "test:replyDataType:INT8:5",  args{ r5, nil } , TypeInt8, false },
		{   "test:replyDataType:INT16:6",  args{ r6, nil } , TypeInt16, false },
		{   "test:replyDataType:INT32:7",  args{ r7, nil } , TypeInt32, false },
		{   "test:replyDataType:INT64:8",  args{ r8, nil } , TypeInt64, false },
		{   "test:replyDataType:UINT8:9",  args{ r9, nil } , TypeUint8, false },
		{   "test:replyDataType:UINT16:10",  args{ r10, nil } , TypeUint16, false },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDt, err := replyDataType(tt.args.reply, tt.args.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("replyDataType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDt != tt.wantDt {
				t.Errorf("replyDataType() gotDt = %v, want %v", gotDt, tt.wantDt)
			}
		})
	}
}