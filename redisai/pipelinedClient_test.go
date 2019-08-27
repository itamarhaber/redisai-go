package redisai

import (
	"github.com/gomodule/redigo/redis"
	"reflect"
	"testing"
)

func TestPipelineIncr(t *testing.T) {
	pclient := ConnectPipelined("redis://localhost:6379", 3, nil)
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
	pclient := ConnectPipelined("redis://localhost:6379", 3, nil)
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
	pclient := ConnectPipelined("redis://localhost:6379", 3, nil)
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
	errorflush := pclient.ForceFlush()
	if errorflush != nil {
		t.Error(errortset)
	}
	if pclient.PipelinePos != 0 {
		t.Errorf("PipelinePos was incorrect, got: %d, want: %d.", pclient.PipelinePos, 0)
	}
	pclient.ActiveConn.Receive()
	rep2, _ := pclient.ActiveConn.Receive()
	dt, shape, data, errProc := ParseTensorResponseValues(rep2)
	if errProc != nil {
		t.Error(errProc)
	}
	if dt != TypeFloat64 {
		t.Errorf("TensorGetValues dt was incorrect, got: %s, want: %s.", dt, TypeFloat64)
	}
	if shape[0] != shp[0] {
		t.Errorf("TensorGetValues shape[0] was incorrect, got: %d, want: %d.", shape[0], shp[0])
	}
	var err error = nil
	dataFloat64s, err := redis.Float64s(data, err)
	if dataFloat64s[0] != values[0] {
		t.Errorf("TensorGetValues dt was incorrect, got: %f, want: %f.", dataFloat64s[0], values[0])
	}
}

func TestConnectPipelined(t *testing.T) {
	type args struct {
		url         string
		pipelineMax int
		pool        *redis.Pool
	}
	tests := []struct {
		name  string
		args  args
		wantC *PipelinedClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := ConnectPipelined(tt.args.url, tt.args.pipelineMax, tt.args.pool); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("ConnectPipelined() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestPipelinedClient_ForceFlush(t *testing.T) {
	type fields struct {
		Pool            *redis.Pool
		PipelineMaxSize int
		PipelinePos     int
		ActiveConn      redis.Conn
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &PipelinedClient{
				Pool:            tt.fields.Pool,
				PipelineMaxSize: tt.fields.PipelineMaxSize,
				PipelinePos:     tt.fields.PipelinePos,
				ActiveConn:      tt.fields.ActiveConn,
			}
			if err := c.ForceFlush(); (err != nil) != tt.wantErr {
				t.Errorf("ForceFlush() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipelinedClient_LoadBackend(t *testing.T) {
	type fields struct {
		Pool            *redis.Pool
		PipelineMaxSize int
		PipelinePos     int
		ActiveConn      redis.Conn
	}
	type args struct {
		backendIdentifier string
		location          string
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
			c := &PipelinedClient{
				Pool:            tt.fields.Pool,
				PipelineMaxSize: tt.fields.PipelineMaxSize,
				PipelinePos:     tt.fields.PipelinePos,
				ActiveConn:      tt.fields.ActiveConn,
			}
			if err := c.LoadBackend(tt.args.backendIdentifier, tt.args.location); (err != nil) != tt.wantErr {
				t.Errorf("LoadBackend() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipelinedClient_ModelDel(t *testing.T) {
	type fields struct {
		Pool            *redis.Pool
		PipelineMaxSize int
		PipelinePos     int
		ActiveConn      redis.Conn
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
			c := &PipelinedClient{
				Pool:            tt.fields.Pool,
				PipelineMaxSize: tt.fields.PipelineMaxSize,
				PipelinePos:     tt.fields.PipelinePos,
				ActiveConn:      tt.fields.ActiveConn,
			}
			if err := c.ModelDel(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("ModelDel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipelinedClient_ModelGet(t *testing.T) {
	type fields struct {
		Pool            *redis.Pool
		PipelineMaxSize int
		PipelinePos     int
		ActiveConn      redis.Conn
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
			c := &PipelinedClient{
				Pool:            tt.fields.Pool,
				PipelineMaxSize: tt.fields.PipelineMaxSize,
				PipelinePos:     tt.fields.PipelinePos,
				ActiveConn:      tt.fields.ActiveConn,
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

func TestPipelinedClient_ModelRun(t *testing.T) {
	type fields struct {
		Pool            *redis.Pool
		PipelineMaxSize int
		PipelinePos     int
		ActiveConn      redis.Conn
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
			c := &PipelinedClient{
				Pool:            tt.fields.Pool,
				PipelineMaxSize: tt.fields.PipelineMaxSize,
				PipelinePos:     tt.fields.PipelinePos,
				ActiveConn:      tt.fields.ActiveConn,
			}
			if err := c.ModelRun(tt.args.name, tt.args.inputs, tt.args.outputs); (err != nil) != tt.wantErr {
				t.Errorf("ModelRun() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipelinedClient_ModelSet(t *testing.T) {
	type fields struct {
		Pool            *redis.Pool
		PipelineMaxSize int
		PipelinePos     int
		ActiveConn      redis.Conn
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
			c := &PipelinedClient{
				Pool:            tt.fields.Pool,
				PipelineMaxSize: tt.fields.PipelineMaxSize,
				PipelinePos:     tt.fields.PipelinePos,
				ActiveConn:      tt.fields.ActiveConn,
			}
			if err := c.ModelSet(tt.args.name, tt.args.backend, tt.args.device, tt.args.data, tt.args.inputs, tt.args.outputs); (err != nil) != tt.wantErr {
				t.Errorf("ModelSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipelinedClient_ScriptDel(t *testing.T) {
	type fields struct {
		Pool            *redis.Pool
		PipelineMaxSize int
		PipelinePos     int
		ActiveConn      redis.Conn
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
			c := &PipelinedClient{
				Pool:            tt.fields.Pool,
				PipelineMaxSize: tt.fields.PipelineMaxSize,
				PipelinePos:     tt.fields.PipelinePos,
				ActiveConn:      tt.fields.ActiveConn,
			}
			if err := c.ScriptDel(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("ScriptDel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipelinedClient_ScriptGet(t *testing.T) {
	type fields struct {
		Pool            *redis.Pool
		PipelineMaxSize int
		PipelinePos     int
		ActiveConn      redis.Conn
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
			c := &PipelinedClient{
				Pool:            tt.fields.Pool,
				PipelineMaxSize: tt.fields.PipelineMaxSize,
				PipelinePos:     tt.fields.PipelinePos,
				ActiveConn:      tt.fields.ActiveConn,
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

func TestPipelinedClient_ScriptRun(t *testing.T) {
	type fields struct {
		Pool            *redis.Pool
		PipelineMaxSize int
		PipelinePos     int
		ActiveConn      redis.Conn
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
			c := &PipelinedClient{
				Pool:            tt.fields.Pool,
				PipelineMaxSize: tt.fields.PipelineMaxSize,
				PipelinePos:     tt.fields.PipelinePos,
				ActiveConn:      tt.fields.ActiveConn,
			}
			if err := c.ScriptRun(tt.args.name, tt.args.fn, tt.args.inputs, tt.args.outputs); (err != nil) != tt.wantErr {
				t.Errorf("ScriptRun() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipelinedClient_ScriptSet(t *testing.T) {
	type fields struct {
		Pool            *redis.Pool
		PipelineMaxSize int
		PipelinePos     int
		ActiveConn      redis.Conn
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
			c := &PipelinedClient{
				Pool:            tt.fields.Pool,
				PipelineMaxSize: tt.fields.PipelineMaxSize,
				PipelinePos:     tt.fields.PipelinePos,
				ActiveConn:      tt.fields.ActiveConn,
			}
			if err := c.ScriptSet(tt.args.name, tt.args.device, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ScriptSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipelinedClient_TensorGet(t *testing.T) {
	type fields struct {
		Pool            *redis.Pool
		PipelineMaxSize int
		PipelinePos     int
		ActiveConn      redis.Conn
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
			c := &PipelinedClient{
				Pool:            tt.fields.Pool,
				PipelineMaxSize: tt.fields.PipelineMaxSize,
				PipelinePos:     tt.fields.PipelinePos,
				ActiveConn:      tt.fields.ActiveConn,
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

func TestPipelinedClient_TensorGetValues(t *testing.T) {
	type fields struct {
		Pool            *redis.Pool
		PipelineMaxSize int
		PipelinePos     int
		ActiveConn      redis.Conn
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
			c := &PipelinedClient{
				Pool:            tt.fields.Pool,
				PipelineMaxSize: tt.fields.PipelineMaxSize,
				PipelinePos:     tt.fields.PipelinePos,
				ActiveConn:      tt.fields.ActiveConn,
			}
			if err := c.TensorGetValues(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("TensorGetValues() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipelinedClient_TensorSet(t *testing.T) {
	type fields struct {
		Pool            *redis.Pool
		PipelineMaxSize int
		PipelinePos     int
		ActiveConn      redis.Conn
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &PipelinedClient{
				Pool:            tt.fields.Pool,
				PipelineMaxSize: tt.fields.PipelineMaxSize,
				PipelinePos:     tt.fields.PipelinePos,
				ActiveConn:      tt.fields.ActiveConn,
			}
			if err := c.TensorSet(tt.args.name, tt.args.dt, tt.args.dims, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("TensorSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipelinedClient_pipeIncr(t *testing.T) {
	type fields struct {
		Pool            *redis.Pool
		PipelineMaxSize int
		PipelinePos     int
		ActiveConn      redis.Conn
	}
	type args struct {
		conn redis.Conn
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
			c := &PipelinedClient{
				Pool:            tt.fields.Pool,
				PipelineMaxSize: tt.fields.PipelineMaxSize,
				PipelinePos:     tt.fields.PipelinePos,
				ActiveConn:      tt.fields.ActiveConn,
			}
			if err := c.pipeIncr(tt.args.conn); (err != nil) != tt.wantErr {
				t.Errorf("pipeIncr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPipelinedClient_Close(t *testing.T) {
	type fields struct {
		Pool            *redis.Pool
		PipelineMaxSize int
		PipelinePos     int
		ActiveConn      redis.Conn
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &PipelinedClient{
				Pool:            tt.fields.Pool,
				PipelineMaxSize: tt.fields.PipelineMaxSize,
				PipelinePos:     tt.fields.PipelinePos,
				ActiveConn:      tt.fields.ActiveConn,
			}
			if err := c.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
