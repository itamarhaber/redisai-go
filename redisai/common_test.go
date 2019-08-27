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

func TestTensorSetArgs(t *testing.T) {
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
		want redis.Args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TensorSetArgs(tt.args.name, tt.args.dt, tt.args.dims, tt.args.data, tt.args.includeCommandName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TensorSetArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
