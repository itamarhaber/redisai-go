package redisai

import (
	"github.com/gomodule/redigo/redis"
	"reflect"
	"testing"
)


// Global vars:
var (
	pclient = Connect("redis://localhost:6379",  nil )
)


func TestClient_LoadBackend(t *testing.T) {
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		backend_identifier BackendType
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

	valuesFloat32 := []float32{1.1}
	valuesFloat64 := []float64{1.1}

	valuesInt8 := []int8{1}
	valuesInt16 := []int16{1}
	valuesInt32 := []int{1}
	valuesInt64 := []int64{1}

	valuesUint8 := []uint8{1}
	valuesUint16 := []uint16{1}
	keyFloat32 := "test:TensorGet:TypeFloat32:1"
	keyFloat64 := "test:TensorGet:TypeFloat64:1"

	keyInt8 := "test:TensorGet:TypeInt8:1"
	keyInt16 := "test:TensorGet:TypeInt16:1"
	keyInt32 := "test:TensorGet:TypeInt32:1"
	keyInt64 := "test:TensorGet:TypeInt64:1"

	keyUint8 := "test:TensorGet:TypeUint8:1"
	keyUint16 := "test:TensorGet:TypeUint16:1"
	shp := []int{1}
	pclient.TensorSet(keyFloat32, TypeFloat32, shp, valuesFloat32)
	pclient.TensorSet(keyFloat64, TypeFloat64, shp, valuesFloat64)

	pclient.TensorSet(keyInt8, TypeInt8, shp, valuesInt8)
	pclient.TensorSet(keyInt16, TypeInt16, shp, valuesInt16)
	pclient.TensorSet(keyInt32, TypeInt32, shp, valuesInt32)
	pclient.TensorSet(keyInt64, TypeInt64, shp, valuesInt64)

	pclient.TensorSet(keyUint8, TypeUint8, shp, valuesUint8)
	pclient.TensorSet(keyUint16, TypeUint16, shp, valuesUint16)

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
		wantDt DataType
		wantShape []int
		wantData interface{}
		compareDt bool
		compareShape bool
		compareData bool
		wantErr  bool
	}{
		{ keyFloat32, fields{ pclient.pool } , args{ keyFloat32, TensorContentTypeValues }, TypeFloat32,shp,valuesFloat32, true, true, true, false},
		{ keyFloat64, fields{ pclient.pool } , args{ keyFloat64, TensorContentTypeValues }, TypeFloat64,shp,valuesFloat64, true, true, true, false},

		{ keyInt8, fields{ pclient.pool } , args{ keyInt8, TensorContentTypeValues }, TypeInt8,shp,valuesInt8, true, true, true, false},
		{ keyInt16, fields{ pclient.pool } , args{ keyInt16, TensorContentTypeValues }, TypeInt16,shp,valuesInt16, true, true, true, false},
		{ keyInt32, fields{ pclient.pool } , args{ keyInt32, TensorContentTypeValues }, TypeInt32,shp,valuesInt32, true, true, true, false},
		{ keyInt64, fields{ pclient.pool } , args{ keyInt64, TensorContentTypeValues }, TypeInt64,shp,valuesInt64, true, true, true, false},
		//commetting while issue is sorted out
		//{ keyUint8, fields{ pclient.pool } , args{ keyUint8, TensorContentTypeValues }, TypeUint8,shp,valuesUint8, true, true, true, false},
		//{ keyUint16, fields{ pclient.pool } , args{ keyUint16, TensorContentTypeValues }, TypeUint16,shp,valuesUint16, true, true, true, false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				pool: tt.fields.pool,
			}
			gotResp, err := c.TensorGet(tt.args.name, tt.args.ct)
			if (err != nil) != tt.wantErr {
				t.Errorf("TensorGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.compareDt && !reflect.DeepEqual(gotResp[0], tt.wantDt) {
				t.Errorf("TensorGet() gotDt = %v, want %v", gotResp[0], tt.wantDt)
			}
			if tt.compareShape && !reflect.DeepEqual(gotResp[1], tt.wantShape) {
				t.Errorf("TensorGet() gotShape = %v, want %v", gotResp[1], tt.wantShape)
			}
			if tt.compareData && !reflect.DeepEqual(gotResp[2], tt.wantData) {
				t.Errorf("TensorGet() gotData = %v, want %v", gotResp[2], tt.wantData)
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

	valuesFloat32 := []float32{1.1}
	valuesFloat64 := []float64{1.1}

	valuesInt8 := []int8{1}
	valuesInt16 := []int16{1}
	valuesInt32 := []int{1}
	valuesInt64 := []int64{1}

	valuesUint8 := []uint8{1}
	valuesUint16 := []uint16{1}
	keyFloat32 := "test:TensorGet:TypeFloat32:1"
	keyFloat64 := "test:TensorGet:TypeFloat64:1"

	keyInt8 := "test:TensorGet:TypeInt8:1"
	keyInt16 := "test:TensorGet:TypeInt16:1"
	keyInt32 := "test:TensorGet:TypeInt32:1"
	keyInt64 := "test:TensorGet:TypeInt64:1"

	keyUint8 := "test:TensorGet:TypeUint8:1"
	keyUint16 := "test:TensorGet:TypeUint16:1"
	shp := []int{1}
	pclient.TensorSet(keyFloat32, TypeFloat32, shp, valuesFloat32)
	pclient.TensorSet(keyFloat64, TypeFloat64, shp, valuesFloat64)

	pclient.TensorSet(keyInt8, TypeInt8, shp, valuesInt8)
	pclient.TensorSet(keyInt16, TypeInt16, shp, valuesInt16)
	pclient.TensorSet(keyInt32, TypeInt32, shp, valuesInt32)
	pclient.TensorSet(keyInt64, TypeInt64, shp, valuesInt64)

	pclient.TensorSet(keyUint8, TypeUint8, shp, valuesUint8)
	pclient.TensorSet(keyUint16, TypeUint16, shp, valuesUint16)

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
		{ keyFloat32, fields{ pclient.pool } , args{ keyFloat32 }, TypeFloat32,shp,valuesFloat32, false},
		{ keyFloat64, fields{ pclient.pool } , args{ keyFloat64 }, TypeFloat64,shp,valuesFloat64,  false},

		{ keyInt8, fields{ pclient.pool } , args{ keyInt8 }, TypeInt8,shp,valuesInt8,false},
		{ keyInt16, fields{ pclient.pool } , args{ keyInt16 }, TypeInt16,shp,valuesInt16, false},
		{ keyInt32, fields{ pclient.pool } , args{ keyInt32 }, TypeInt32,shp,valuesInt32,  false},
		{ keyInt64, fields{ pclient.pool } , args{ keyInt64 }, TypeInt64,shp,valuesInt64,  false},

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

	valuesFloat32 := []float32{1.1}
	valuesFloat64 := []float64{1.1}

	valuesInt8 := []int8{1}
	valuesInt16 := []int16{1}
	valuesInt32 := []int{1}
	valuesInt64 := []int64{1}

	valuesUint8 := []uint8{1}
	valuesByte := []byte{1}

	valuesUint16 := []uint16{1}
	keyFloat32 := "test:TensorSet:TypeFloat32:1"
	keyFloat64 := "test:TensorSet:TypeFloat64:1"

	keyInt8 := "test:TensorSet:TypeInt8:1"
	keyInt16 := "test:TensorSet:TypeInt16:1"
	keyInt32 := "test:TensorSet:TypeInt32:1"
	keyInt64 := "test:TensorSet:TypeInt64:1"

	keyByte := "test:TensorSet:Type[]byte:1"
	keyUint8 := "test:TensorSet:TypeUint8:1"
	keyUint16 := "test:TensorSet:TypeUint16:1"

	keyInt8Meta := "test:TensorSet:TypeInt8:Meta:1"

	shp := []int{1}

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
		{ keyFloat32, fields{ pclient.pool } , args{ keyFloat32, TypeFloat, shp, valuesFloat32 }, false },
		{ keyFloat64, fields{ pclient.pool } , args{ keyFloat64, TypeFloat64, shp, valuesFloat64 }, false },

		{ keyInt8, fields{ pclient.pool } , args{ keyInt8, TypeInt8, shp, valuesInt8 }, false },
		{ keyInt16, fields{ pclient.pool } , args{ keyInt16, TypeInt16, shp, valuesInt16 }, false },
		{ keyInt32, fields{ pclient.pool } , args{ keyInt32, TypeInt32, shp, valuesInt32 }, false },
		{ keyInt64, fields{ pclient.pool } , args{ keyInt64, TypeInt64, shp, valuesInt64 }, false },

		{ keyUint8, fields{ pclient.pool } , args{ keyUint8, TypeUint8, shp, valuesUint8 }, false },
		{ keyUint16, fields{ pclient.pool } , args{ keyUint16, TypeUint16, shp, valuesUint16 }, false },

		{ keyInt8Meta, fields{ pclient.pool } , args{ keyInt8Meta, TypeUint8, shp, nil }, false },
		{ keyByte, fields{ pclient.pool } , args{ keyByte, TypeUint8, shp, valuesByte }, false },

		{ "test:TestClient_TensorSet:1:FaultyDims", fields{ pclient.pool } , args{ "test:TestClient_TensorSet:1:FaultyDims", TypeFloat, []int{1,10}, []float32{1} }, true },
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
