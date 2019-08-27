package redisai

import (
	"github.com/gomodule/redigo/redis"
	"reflect"
	"testing"
)

func TestClient_LoadBackend(t *testing.T) {
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		backend_identifier string
		location           string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pool: tt.fields.pool,
			}
			if err := c.LoadBackend(tt.args.backend_identifier, tt.args.location); (err != nil) != tt.wantErr {
				t.Errorf("LoadBackend() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_ModelDel(t *testing.T) {
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pool: tt.fields.pool,
			}
			if err := c.ModelDel(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("ModelDel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_ModelGet(t *testing.T) {
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		name string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData []byte
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pool: tt.fields.pool,
			}
			gotData, err := c.ModelGet(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModelGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("ModelGet() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestClient_ModelRun(t *testing.T) {
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		name    string
		inputs  []string
		outputs []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pool: tt.fields.pool,
			}
			if err := c.ModelRun(tt.args.name, tt.args.inputs, tt.args.outputs); (err != nil) != tt.wantErr {
				t.Errorf("ModelRun() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_ModelSet(t *testing.T) {
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		name    string
		backend BackendType
		device  DeviceType
		data    []byte
		inputs  []string
		outputs []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pool: tt.fields.pool,
			}
			if err := c.ModelSet(tt.args.name, tt.args.backend, tt.args.device, tt.args.data, tt.args.inputs, tt.args.outputs); (err != nil) != tt.wantErr {
				t.Errorf("ModelSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_ModelSetFromFile(t *testing.T) {
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		name    string
		backend BackendType
		device  DeviceType
		path    string
		inputs  []string
		outputs []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pool: tt.fields.pool,
			}
			if err := c.ModelSetFromFile(tt.args.name, tt.args.backend, tt.args.device, tt.args.path, tt.args.inputs, tt.args.outputs); (err != nil) != tt.wantErr {
				t.Errorf("ModelSetFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_ScriptDel(t *testing.T) {
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pool: tt.fields.pool,
			}
			if err := c.ScriptDel(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("ScriptDel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_ScriptGet(t *testing.T) {
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		name string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData []byte
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pool: tt.fields.pool,
			}
			gotData, err := c.ScriptGet(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScriptGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("ScriptGet() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestClient_ScriptRun(t *testing.T) {
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		name    string
		fn      string
		inputs  []string
		outputs []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pool: tt.fields.pool,
			}
			if err := c.ScriptRun(tt.args.name, tt.args.fn, tt.args.inputs, tt.args.outputs); (err != nil) != tt.wantErr {
				t.Errorf("ScriptRun() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_ScriptSet(t *testing.T) {
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		name   string
		device DeviceType
		data   []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pool: tt.fields.pool,
			}
			if err := c.ScriptSet(tt.args.name, tt.args.device, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ScriptSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_ScriptSetFromFile(t *testing.T) {
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		name   string
		device DeviceType
		path   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pool: tt.fields.pool,
			}
			if err := c.ScriptSetFromFile(tt.args.name, tt.args.device, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("ScriptSetFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_TensorGet(t *testing.T) {
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		name string
		ct   TensorContentType
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData interface{}
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pool: tt.fields.pool,
			}
			gotData, err := c.TensorGet(tt.args.name, tt.args.ct)
			if (err != nil) != tt.wantErr {
				t.Errorf("TensorGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("TensorGet() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestClient_TensorGetBlob(t *testing.T) {
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		fields    fields
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
			c := &Client{
				pool: tt.fields.pool,
			}
			gotDt, gotShape, gotData, err := c.TensorGetBlob(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("TensorGetBlob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDt != tt.wantDt {
				t.Errorf("TensorGetBlob() gotDt = %v, want %v", gotDt, tt.wantDt)
			}
			if !reflect.DeepEqual(gotShape, tt.wantShape) {
				t.Errorf("TensorGetBlob() gotShape = %v, want %v", gotShape, tt.wantShape)
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("TensorGetBlob() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestClient_TensorGetMeta(t *testing.T) {
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantDt    DataType
		wantShape []int
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pool: tt.fields.pool,
			}
			gotDt, gotShape, err := c.TensorGetMeta(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("TensorGetMeta() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDt != tt.wantDt {
				t.Errorf("TensorGetMeta() gotDt = %v, want %v", gotDt, tt.wantDt)
			}
			if !reflect.DeepEqual(gotShape, tt.wantShape) {
				t.Errorf("TensorGetMeta() gotShape = %v, want %v", gotShape, tt.wantShape)
			}
		})
	}
}

func TestClient_TensorGetValues(t *testing.T) {
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		fields    fields
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
			c := &Client{
				pool: tt.fields.pool,
			}
			gotDt, gotShape, gotData, err := c.TensorGetValues(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("TensorGetValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDt != tt.wantDt {
				t.Errorf("TensorGetValues() gotDt = %v, want %v", gotDt, tt.wantDt)
			}
			if !reflect.DeepEqual(gotShape, tt.wantShape) {
				t.Errorf("TensorGetValues() gotShape = %v, want %v", gotShape, tt.wantShape)
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("TensorGetValues() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestClient_TensorSet(t *testing.T) {
	pclient := Connect("redis://localhost:6379",  nil )
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		name string
		dt   DataType
		dims []int
		data interface{}
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{ "test:TestPipelineIncr:1", fields{ pclient.pool } , args{ "test:TestPipelineIncr:1", TypeFloat, []int{1}, []float32{1} }, false },
		{ "test:TestPipelineIncr:1:FaultyDims", fields{ pclient.pool } , args{ "test:TestPipelineIncr:1", TypeFloat, []int{1,10}, []float32{1} }, true },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pool: tt.fields.pool,
			}
			if err := c.TensorSet(tt.args.name, tt.args.dt, tt.args.dims, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("TensorSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConnect(t *testing.T) {
	type args struct {
		url  string
		pool *redis.Pool
	}
	tests := []struct {
		name  string
		args  args
		wantC *Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := Connect(tt.args.url, tt.args.pool); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("Connect() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
