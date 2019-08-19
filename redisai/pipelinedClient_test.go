package redisai

import (
	"testing"
)

func TestPipelineIncr(t *testing.T) {
	pclient := ConnectPipelined("redis://localhost:6379", 3)
	defer pclient.Close()
	errortset := pclient.TensorSet("test:TestPipelineIncr:1", TypeFloat, []int{1}, []float32{1})
	if errortset != nil {
		t.Error(errortset)
	}
	if pclient.PipelinePos != 1 {
		t.Errorf("PipelinePos was incorrect, got: %d, want: %d.", pclient.PipelinePos, 1)
	}
}

func TestPipelineResetOnLimit(t *testing.T) {
	pclient := ConnectPipelined("redis://localhost:6379", 3)
	defer pclient.Close()
	errortset := pclient.TensorSet("test:TestPipelineResetOnLimit:1", TypeFloat, []int{4}, []float32{1.1, 2.2, 3.3, 4.4})
	if errortset != nil {
		t.Error(errortset)
	}
	if pclient.PipelinePos != 1 {
		t.Errorf("PipelinePos was incorrect, got: %d, want: %d.", pclient.PipelinePos, 1)
	}
	errortset = pclient.TensorSet("test:TestPipelineResetOnLimit:2", TypeFloat, []int{1, 2}, []float32{1.1, 2.2})
	if errortset != nil {
		t.Error(errortset)
	}
	if pclient.PipelinePos != 2 {
		t.Errorf("PipelinePos was incorrect, got: %d, want: %d.", pclient.PipelinePos, 2)
	}
	errortset = pclient.TensorSet("test:TestPipelineResetOnLimit:3", TypeFloat, []int{1, 3}, []float32{1.1, 2.2, 3.3})
	if errortset != nil {
		t.Error(errortset)
	}

	if pclient.PipelinePos != 0 {
		t.Errorf("PipelinePos was incorrect, got: %d, want: %d.", pclient.PipelinePos, 0)
	}
}


func TestTensorGetValues(t *testing.T) {
	pclient := ConnectPipelined("redis://localhost:6379", 3)
	values := []float64{1.1, 2.2, 3.3, 4.4}
	shp := []int{4}
	defer pclient.Close()
	errortset := pclient.TensorSet("test:TensorGetValues:tensor1", TypeFloat64, shp, values)
	if errortset != nil {
		t.Error(errortset)
	}
	if pclient.PipelinePos != 1 {
		t.Errorf("PipelinePos was incorrect, got: %d, want: %d.", pclient.PipelinePos, 1)
	}
	errortget := pclient.TensorGetValues("test:TensorGetValues:tensor1")
	if errortget != nil {
		t.Error(errortget)
	}
	if pclient.PipelinePos != 2 {
		t.Errorf("PipelinePos was incorrect, got: %d, want: %d.", pclient.PipelinePos, 2)
	}
	errorflush:= pclient.forceFlush()
	if errorflush != nil {
		t.Error(errortset)
	}
	if pclient.PipelinePos != 0 {
		t.Errorf("PipelinePos was incorrect, got: %d, want: %d.", pclient.PipelinePos, 0)
	}
	 pclient.ActiveConn.Receive()
	rep2 , _ := pclient.ActiveConn.Receive()
	dt, shape, data, errProc := ProcessTensorResponse(rep2)
	if errProc != nil {
		t.Error(errProc)
	}
	if dt != TypeFloat64 {
		t.Errorf("TensorGetValues dt was incorrect, got: %s, want: %s.", dt, TypeFloat64)
	}
	if shape[0] != shp[0] {
		t.Errorf("TensorGetValues shape[0] was incorrect, got: %d, want: %d.", shape[0], shp[0])
	}
	if data[0] != values[0] {
		t.Errorf("TensorGetValues dt was incorrect, got: %f, want: %f.", data[0], values[0])
	}

}

func TestTensorSetArgs(t *testing.T) {

}
