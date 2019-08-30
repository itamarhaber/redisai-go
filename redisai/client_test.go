package redisai

import (
	"github.com/gomodule/redigo/redis"
	"io/ioutil"
	"reflect"
	"testing"
)

// Global vars:
var (
	pclient = Connect("redis://localhost:6379", nil)
)

func TestClient_LoadBackend(t *testing.T) {
	keyTest1 := "test:LoadBackend:Unexistant:1"
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
		{keyTest1, fields{pclient.pool}, args{BackendTF, "unexistant"}, true},

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
	keyModel1 := "test:ModelDel:1"
	data, err := ioutil.ReadFile("./../tests/testdata/models/tensorflow/creditcardfraud.pb")
	if err != nil {
		t.Errorf("Error preparing for ModelDel(), while issuing ModelSet. error = %v", err)
		return
	}
	err = pclient.ModelSet(keyModel1, BackendTF, DeviceCPU, data, []string{"transaction", "reference"}, []string{"output"})
	if err != nil {
		t.Errorf("Error preparing for ModelDel(), while issuing ModelSet. error = %v", err)
		return
	}
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
		{keyModel1, fields{pclient.pool}, args{keyModel1}, false},
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
	keyModel1 := "test:ModelGet:1"
	keyModelUnexistent1 := "test:ModelGetUnexistent:1"
	data, err := ioutil.ReadFile("./../tests/testdata/models/tensorflow/creditcardfraud.pb")
	if err != nil {
		t.Errorf("Error preparing for ModelGet(), while issuing ModelSet. error = %v", err)
		return
	}
	err = pclient.ModelSet(keyModel1, BackendTF, DeviceCPU, data, []string{"transaction", "reference"}, []string{"output"})
	if err != nil {
		t.Errorf("Error preparing for ModelGet(), while issuing ModelSet. error = %v", err)
		return
	}
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		name string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantBackend BackendType
		wantDevice  DeviceType
		wantData    []byte
		wantErr     bool
		testBackend bool
		testDevice  bool
		testData    bool
	}{
		{keyModelUnexistent1, fields{pclient.pool}, args{keyModelUnexistent1}, BackendTF, DeviceCPU, data, true, false, false, false},
		{keyModel1, fields{pclient.pool}, args{keyModel1}, BackendTF, DeviceCPU, data, false, true, true, false},
		// TODO: check why is failing
		//{ keyModel1, fields{ pclient.pool } , args{ keyModel1 }, BackendTF, DeviceCPU ,data ,false,true,true,true},
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
			if tt.testBackend {
				if !reflect.DeepEqual(gotData[0], tt.wantBackend) {
					t.Errorf("ModelGet() gotBackend = %v, want %v. gotBackend Type %v, want Type %v.", gotData[0], tt.wantBackend, reflect.TypeOf(gotData[0]), reflect.TypeOf(tt.wantBackend))
				}
			}
			if tt.testDevice {
				if !reflect.DeepEqual(gotData[1], tt.wantDevice) {
					t.Errorf("ModelGet() gotDevice = %v, want %v. gotDevice Type %v, want Type %v.", gotData[1], tt.wantDevice, reflect.TypeOf(gotData[1]), reflect.TypeOf(tt.wantDevice))
				}
			}
			if tt.testData {
				if !reflect.DeepEqual(gotData[2], tt.wantData) {
					t.Errorf("ModelGet() gotData = %v, want %v. gotData Type %v, want Type %v.", gotData[2], tt.wantData, reflect.TypeOf(gotData[2]), reflect.TypeOf(tt.wantData))
				}
			}

		})
	}
}

func TestClient_ModelRun(t *testing.T) {
	keyModel1 := "test:ModelRun:1"
	keyModelWrongInput1 := "test:ModelWrongInput:1"
	keyTransaction1 := "test:ModelRun:transaction:1"
	keyReference1 := "test:ModelRun:reference:1"
	keyOutput1 := "test:ModelRun:output:1"

	data, err := ioutil.ReadFile("./../tests/testdata/models/tensorflow/creditcardfraud.pb")
	if err != nil {
		t.Errorf("Error preparing for ModelRun(), while issuing ModelSet. error = %v", err)
		return
	}
	err = pclient.ModelSet(keyModel1, BackendTF, DeviceCPU, data, []string{"transaction", "reference"}, []string{"output"})
	if err != nil {
		t.Errorf("Error preparing for ModelRun(), while issuing ModelSet. error = %v", err)
		return
	}

	errortset := pclient.TensorSet(keyTransaction1, TypeFloat, []int{1, 30}, []float32{0,
		-1.3598071336738,
		-0.0727811733098497,
		2.53634673796914,
		1.37815522427443,
		-0.338320769942518,
		0.462387777762292,
		0.239598554061257,
		0.0986979012610507,
		0.363786969611213,
		0.0907941719789316,
		-0.551599533260813,
		-0.617800855762348,
		-0.991389847235408,
		-0.311169353699879,
		1.46817697209427,
		-0.470400525259478,
		0.207971241929242,
		0.0257905801985591,
		0.403992960255733,
		0.251412098239705,
		-0.018306777944153,
		0.277837575558899,
		-0.110473910188767,
		0.0669280749146731,
		0.128539358273528,
		-0.189114843888824,
		0.133558376740387,
		-0.0210530534538215,
		149.62})
	if errortset != nil {
		t.Error(errortset)
	}

	errortsetReference := pclient.TensorSet(keyReference1, TypeFloat, []int{256}, []float32{
		0.4961020023739511,
		0.25008885268782743,
		0.17356637875650527,
		0.455499134765027,
		0.36590153405192427,
		0.025643573760428695,
		0.911348673549787,
		0.280600196316659,
		0.19903348845122615,
		0.4843237748033392,
		0.0466782567080819,
		0.39655339475845763,
		0.08886225131143377,
		0.8975580119246835,
		0.5876046939685196,
		0.2036572605491107,
		0.49587805570111154,
		0.7153861813728742,
		0.9156194444373905,
		0.2502921311442605,
		0.8048228543655253,
		0.3786155916087869,
		0.24695783264314564,
		0.9634375461649354,
		0.6347336474822765,
		0.625234717098543,
		0.10027243263221086,
		0.5152389510603593,
		0.24729154458831293,
		0.2017559178166548,
		0.93168739414145,
		0.20110380520573967,
		0.31179378782980205,
		0.21000262832227234,
		0.7364270692603087,
		0.19993210868657152,
		0.7318737388858069,
		0.20875355445773913,
		0.445224688232584,
		0.9774779314791744,
		0.5326116359851079,
		0.5101212498135284,
		0.7086788842415588,
		0.6147374814798513,
		0.2016813834414265,
		0.409418198470738,
		0.8255491375035944,
		0.6786705045501186,
		0.7236519487406021,
		0.10979804308949248,
		0.9477441831989238,
		0.45719805166675387,
		0.610153730100084,
		0.11655669231561605,
		0.4439894014709225,
		0.7446443906737652,
		0.8216651981976272,
		0.5789391572965281,
		0.014031859961184279,
		0.4520095606042871,
		0.9825890727240326,
		0.7886484063650101,
		0.77627752119412,
		0.4481386679813363,
		0.8793965874947762,
		0.6917286714705064,
		0.6856714599658206,
		0.5935835205953005,
		0.40373465761470895,
		0.4103196001041468,
		0.8466047746635962,
		0.12585140814309892,
		0.1275895372313478,
		0.36862564073917303,
		0.7909646987305703,
		0.6535528917624062,
		0.2944289897532757,
		0.9230888632644605,
		0.30215077639978694,
		0.7104415296747062,
		0.23358534963067223,
		0.20267464409166136,
		0.43805968728761757,
		0.1360918122594953,
		0.9603124922591536,
		0.09217517262849939,
		0.3934965742783815,
		0.9880379118731525,
		0.4157802751771462,
		0.36351834248258585,
		0.203470028463675,
		0.5644122076867265,
		0.3607003042390333,
		0.8627479960712836,
		0.896717617812036,
		0.07194770994261201,
		0.40360859469525656,
		0.710234618370674,
		0.39669402322777003,
		0.7588202069029378,
		0.3967493109500312,
		0.971726089964839,
		0.09743562226055202,
		0.24826374660523043,
		0.6555458927354575,
		0.5471342964153852,
		0.2459704064955166,
		0.6262311367955701,
		0.3344751806822718,
		0.22114039088261161,
		0.9923586561385392,
		0.24894482650730698,
		0.6327454030779037,
		0.25354887978857366,
		0.5478295365352684,
		0.07989786035960178,
		0.5440351740551437,
		0.012914207986969628,
		0.46727537723784385,
		0.6590735810404428,
		0.9389135387540076,
		0.31765723308475124,
		0.5937715350874003,
		0.7172974278007461,
		0.3955878908785877,
		0.06712667047697007,
		0.39789421966780925,
		0.08840426188349138,
		0.6288925675386916,
		0.27112782019946136,
		0.4116628783835494,
		0.13365791270780514,
		0.4864959836599304,
		0.835891040614729,
		0.46417516300140194,
		0.7513645163836994,
		0.4919312892675719,
		0.8785225152156605,
		0.5525317575031543,
		0.3918884347804765,
		0.48070860728006914,
		0.3323215096874963,
		0.7456924987916765,
		0.3845226328482302,
		0.41184851469429595,
		0.4970158291960127,
		0.15085629972627568,
		0.21903528393808147,
		0.23057635019441947,
		0.09509620966166554,
		0.8605106738443453,
		0.3382348342856798,
		0.5462936342674528,
		0.6197259060274013,
		0.45400416154184053,
		0.43153012489457965,
		0.9598194428132951,
		0.41465122328276816,
		0.4698336388751333,
		0.8407476256896753,
		0.8991897604039162,
		0.5871733369659597,
		0.9489727535807733,
		0.03966682159646773,
		0.059638838923675386,
		0.6480914849839939,
		0.0032055103028040266,
		0.5644179356077625,
		0.6237238941355112,
		0.7426357772990153,
		0.48708641552158627,
		0.5738652541551791,
		0.399452394520291,
		0.11315150790074868,
		0.4463757751498464,
		0.3491631084369967,
		0.7155340294057289,
		0.5486828325884815,
		0.027936967943904878,
		0.6247855870250584,
		0.07760076108013958,
		0.49931433545416615,
		0.5506092158753837,
		0.9943035848277743,
		0.20573445846451,
		0.7216285512945004,
		0.09516133459004761,
		0.8403506939931851,
		0.10933786589888539,
		0.4443788740790786,
		0.470057979424499,
		0.8780383573192116,
		0.8689890461906095,
		0.10756346192407429,
		0.5782064960219897,
		0.6881089157793148,
		0.9105474107882497,
		0.29221759114939505,
		0.3094779191891116,
		0.19817046920128678,
		0.3459723279441753,
		0.985513249223403,
		0.6317298309892471,
		0.10494448511804233,
		0.09885467433452855,
		0.3962644530139615,
		0.29570548319787604,
		0.9509755871106149,
		0.3841769458302071,
		0.26240807237752084,
		0.5243268350123865,
		0.5454667676472065,
		0.03202596453912032,
		0.3139580685666843,
		0.7316746334330743,
		0.01773037472929495,
		0.9693262316508454,
		0.43479823811937035,
		0.05605391232132073,
		0.6563470571241352,
		0.802771930231375,
		0.982625623283999,
		0.7634709307919724,
		0.6821161791991082,
		0.7562380433686934,
		0.6857467886014096,
		0.04303926774340816,
		0.32833800470114494,
		0.042278653840651326,
		0.6569849279597196,
		0.2986179861617654,
		0.47636816550296346,
		0.9864885302198588,
		0.10321289993582661,
		0.4599323272874871,
		0.1925191713379878,
		0.16558404193162335,
		0.36954765643996235,
		0.053636651883167796,
		0.5626652817821394,
		0.859443082864307,
		0.48767732197736513,
		0.4766202894660617,
		0.45643935942565717,
		0.5846504463655596,
		0.12611650476119274,
		0.866971601546102,
		0.4723266255234033,
		0.06573755550521643,
		0.27551870358508623,
		0.08954775454593156,
		0.2102171158669528,
		0.9350683386165229,
		0.4302997612039827,
		0.30431237188340277,
		0.24787823194966807})
	if errortsetReference != nil {
		t.Error(errortsetReference)
	}

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
		{keyModel1, fields{pclient.pool}, args{keyModel1, []string{keyTransaction1, keyReference1}, []string{keyOutput1}}, false},
		{keyModelWrongInput1, fields{pclient.pool}, args{keyModel1, []string{keyTransaction1}, []string{keyOutput1}}, true},
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

	keyModelSet1 := "test:ModelSet:1"
	keyModelSetUnexistant := "test:ModelSet:Unexistant:1"
	dataUnexistant := []byte{}
	data, err := ioutil.ReadFile("./../tests/testdata/models/tensorflow/creditcardfraud.pb")
	if err != nil {
		t.Errorf("Error preparing for ModelSet(), while reading file. error = %v", err)
		return
	}

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
		{keyModelSet1, fields{pclient.pool}, args{keyModelSet1, BackendTF, DeviceCPU, data, []string{"transaction", "reference"}, []string{"output"}}, false},
		{keyModelSetUnexistant, fields{pclient.pool}, args{keyModelSetUnexistant, BackendTF, DeviceCPU, dataUnexistant, []string{"transaction", "reference"}, []string{"output"}}, true},
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
	keyModel1 := "test:ModelSetFromFile:1"
	keyModelDontExist1 := "test:ModelSetFromFile:DontExist:1"
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
		{keyModelDontExist1, fields{pclient.pool}, args{keyModelDontExist1, BackendTF, DeviceCPU, "./../tests/testdata/dontexist", []string{"transaction", "reference"}, []string{"output"}}, true},
		{keyModel1, fields{pclient.pool}, args{keyModel1, BackendTF, DeviceCPU, "./../tests/testdata/models/tensorflow/creditcardfraud.pb", []string{"transaction", "reference"}, []string{"output"}}, false},
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
	keyScript := "test:ScriptDel:1"
	keyScriptUnexistant := "test:ScriptDel:Unexistant:1"
	scriptBin := "def bar(a, b):\n    return a + b\n"

	err := pclient.ScriptSet(keyScript, DeviceCPU, scriptBin)
	if err != nil {
		t.Errorf("Error preparing for ScriptDel(), while issuing ScriptSet. error = %v", err)
		return
	}
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
		{keyScript, fields{pclient.pool}, args{keyScript}, false},
		{keyScriptUnexistant, fields{pclient.pool}, args{keyScriptUnexistant}, true},
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
	keyScript := "test:ScriptGet:1"
	keyScriptEmpty := "test:ScriptGetEmpty:1"

	scriptBin := "def bar(a, b):\n    return a + b\n"

	err := pclient.ScriptSet(keyScript, DeviceCPU, scriptBin)
	if err != nil {
		t.Errorf("Error preparing for ScriptGet(), while issuing ScriptSet. error = %v", err)
		return
	}
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
		wantData []interface{}
		wantErr  bool
	}{
		//TODO: revise this
		//{ keyScript, fields{ pclient.pool } , args{ keyScript }, scriptBin ,false},
		{keyScriptEmpty, fields{pclient.pool}, args{keyScriptEmpty}, nil, true},
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

	keyScript := "test:ScriptRun:1"
	keyScriptEmpty := "test:ScriptRunEmpty:1"

	scriptBin := "def bar(a, b):\n    return a + b\n"

	err := pclient.ScriptSet(keyScript, DeviceCPU, scriptBin)
	if err != nil {
		t.Errorf("Error preparing for ScriptRun(), while issuing ScriptSet. error = %v", err)
		return
	}

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
		{keyScriptEmpty, fields{pclient.pool}, args{keyScriptEmpty, "", []string{""}, []string{""}}, true},
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
	keyScriptError := "test:ScriptSet:Error:1"
	scriptBin := "import abc"
	type fields struct {
		pool *redis.Pool
	}
	type args struct {
		name   string
		device DeviceType
		data   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{keyScriptError, fields{pclient.pool}, args{keyScriptError, DeviceCPU, scriptBin}, true},
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
	keyScript1 := "test:ScriptSetFromFile:DontExist:1"
	keyScript2 := "test:ScriptSetFromFile:2"

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
		{keyScript1, fields{pclient.pool}, args{keyScript1, DeviceCPU, "./../tests/testdata/dontexist"}, true},
		{keyScript2, fields{pclient.pool}, args{keyScript2, DeviceCPU, "./../tests/testdata/script.txt"}, false},
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
	valuesByteSlice := []byte{1, 2, 3, 4, 5}

	valuesFloat32 := []float32{1.1}
	valuesFloat64 := []float64{1.1}

	valuesInt8 := []int8{1}
	valuesInt16 := []int16{1}
	valuesInt32 := []int{1}
	valuesInt64 := []int64{1}

	valuesUint8 := []uint8{1}
	valuesUint16 := []uint16{1}

	keyByteSlice := "test:TensorGet:[]byte:1"
	keyFloat32 := "test:TensorGet:TypeFloat32:1"
	keyFloat64 := "test:TensorGet:TypeFloat64:1"

	keyInt8 := "test:TensorGet:TypeInt8:1"
	keyInt16 := "test:TensorGet:TypeInt16:1"
	keyInt32 := "test:TensorGet:TypeInt32:1"
	keyInt64 := "test:TensorGet:TypeInt64:1"

	keyUint8 := "test:TensorGet:TypeUint8:1"
	keyUint16 := "test:TensorGet:TypeUint16:1"
	shp := []int{1}
	shpByteSlice := []int{1, 5}
	pclient.TensorSet(keyByteSlice, TypeUint8, shpByteSlice, valuesByteSlice)

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
		name         string
		fields       fields
		args         args
		wantDt       DataType
		wantShape    []int
		wantData     interface{}
		compareDt    bool
		compareShape bool
		compareData  bool
		wantErr      bool
	}{
		{keyByteSlice, fields{pclient.pool}, args{keyByteSlice, TensorContentTypeBlob}, TypeUint8, shpByteSlice, valuesByteSlice, true, true, true, false},

		{keyFloat32, fields{pclient.pool}, args{keyFloat32, TensorContentTypeValues}, TypeFloat32, shp, valuesFloat32, true, true, true, false},
		{keyFloat64, fields{pclient.pool}, args{keyFloat64, TensorContentTypeValues}, TypeFloat64, shp, valuesFloat64, true, true, true, false},

		{keyInt8, fields{pclient.pool}, args{keyInt8, TensorContentTypeValues}, TypeInt8, shp, valuesInt8, true, true, true, false},
		{keyInt16, fields{pclient.pool}, args{keyInt16, TensorContentTypeValues}, TypeInt16, shp, valuesInt16, true, true, true, false},
		{keyInt32, fields{pclient.pool}, args{keyInt32, TensorContentTypeValues}, TypeInt32, shp, valuesInt32, true, true, true, false},
		{keyInt64, fields{pclient.pool}, args{keyInt64, TensorContentTypeValues}, TypeInt64, shp, valuesInt64, true, true, true, false},

		{keyUint8, fields{pclient.pool}, args{keyUint8, TensorContentTypeValues}, TypeUint8, shp, valuesUint8, true, true, true, false},
		{keyUint16, fields{pclient.pool}, args{keyUint16, TensorContentTypeValues}, TypeUint16, shp, valuesUint16, true, true, true, false},
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
	valuesByte := []byte{1, 2, 3, 4}
	keyByte := "test:TensorGetBlog:[]byte:1"
	keyUnexistant := "test:TensorGetMeta:Unexistant"

	shp := []int{1, 4}
	pclient.TensorSet(keyByte, TypeInt8, shp, valuesByte)

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
		{keyByte, fields{pclient.pool}, args{keyByte}, TypeInt8, shp, valuesByte, false},
		{keyUnexistant, fields{pclient.pool}, args{keyUnexistant}, TypeInt8, shp, valuesByte, true},
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

			if tt.wantErr == false {
				if gotDt != tt.wantDt {
					t.Errorf("TensorGetBlob() gotDt = %v, want %v", gotDt, tt.wantDt)
				}
				if !reflect.DeepEqual(gotShape, tt.wantShape) {
					t.Errorf("TensorGetBlob() gotShape = %v, want %v", gotShape, tt.wantShape)
				}
				if !reflect.DeepEqual(gotData, tt.wantData) {
					t.Errorf("TensorGetBlob() gotData = %v, want %v", gotData, tt.wantData)
				}
			}
		})
	}
}

func TestClient_TensorGetMeta(t *testing.T) {

	keyFloat32 := "test:TensorGetMeta:TypeFloat32:1"
	keyFloat64 := "test:TensorGetMeta:TypeFloat64:1"

	keyInt8 := "test:TensorGetMeta:TypeInt8:1"
	keyInt16 := "test:TensorGetMeta:TypeInt16:1"
	keyInt32 := "test:TensorGetMeta:TypeInt32:1"
	keyInt64 := "test:TensorGetMeta:TypeInt64:1"

	keyUint8 := "test:TensorGetMeta:TypeUint8:1"
	keyUint16 := "test:TensorGetMeta:TypeUint16:1"

	keyUnexistant := "test:TensorGetMeta:Unexistant"

	shp := []int{1, 2}

	pclient.TensorSet(keyFloat32, TypeFloat32, shp, nil)
	pclient.TensorSet(keyFloat64, TypeFloat64, shp, nil)

	pclient.TensorSet(keyInt8, TypeInt8, shp, nil)
	pclient.TensorSet(keyInt16, TypeInt16, shp, nil)
	pclient.TensorSet(keyInt32, TypeInt32, shp, nil)
	pclient.TensorSet(keyInt64, TypeInt64, shp, nil)

	pclient.TensorSet(keyUint8, TypeUint8, shp, nil)
	pclient.TensorSet(keyUint16, TypeUint16, shp, nil)

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
		{keyFloat32, fields{pclient.pool}, args{keyFloat32}, TypeFloat32, shp, false},
		{keyFloat64, fields{pclient.pool}, args{keyFloat64}, TypeFloat64, shp, false},
		{keyInt8, fields{pclient.pool}, args{keyInt8}, TypeInt8, shp, false},
		{keyInt16, fields{pclient.pool}, args{keyInt16}, TypeInt16, shp, false},
		{keyInt32, fields{pclient.pool}, args{keyInt32}, TypeInt32, shp, false},
		{keyInt64, fields{pclient.pool}, args{keyInt64}, TypeInt64, shp, false},
		{keyUnexistant, fields{pclient.pool}, args{keyUnexistant}, TypeInt64, shp, true},
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
			if tt.wantErr == false {
				if gotDt != tt.wantDt {
					t.Errorf("TensorGetMeta() gotDt = %v, want %v", gotDt, tt.wantDt)
				}
				if !reflect.DeepEqual(gotShape, tt.wantShape) {
					t.Errorf("TensorGetMeta() gotShape = %v, want %v", gotShape, tt.wantShape)
				}
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
	keyFloat32 := "test:TensorGetValues:TypeFloat32:1"
	keyFloat64 := "test:TensorGetValues:TypeFloat64:1"

	keyInt8 := "test:TensorGetValues:TypeInt8:1"
	keyInt16 := "test:TensorGetValues:TypeInt16:1"
	keyInt32 := "test:TensorGetValues:TypeInt32:1"
	keyInt64 := "test:TensorGetValues:TypeInt64:1"

	keyUint8 := "test:TensorGetValues:TypeUint8:1"
	keyUint16 := "test:TensorGetValues:TypeUint16:1"
	keyUnexistant := "test:TensorGetValues:Unexistant"

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
		{keyFloat32, fields{pclient.pool}, args{keyFloat32}, TypeFloat32, shp, valuesFloat32, false},
		{keyFloat64, fields{pclient.pool}, args{keyFloat64}, TypeFloat64, shp, valuesFloat64, false},

		{keyInt8, fields{pclient.pool}, args{keyInt8}, TypeInt8, shp, valuesInt8, false},
		{keyInt16, fields{pclient.pool}, args{keyInt16}, TypeInt16, shp, valuesInt16, false},
		{keyInt32, fields{pclient.pool}, args{keyInt32}, TypeInt32, shp, valuesInt32, false},
		{keyInt64, fields{pclient.pool}, args{keyInt64}, TypeInt64, shp, valuesInt64, false},

		{keyUint8, fields{pclient.pool}, args{keyUint8}, TypeUint8, shp, valuesUint8, false},
		{keyUint16, fields{pclient.pool}, args{keyUint16}, TypeUint16, shp, valuesUint16, false},

		{keyUnexistant, fields{pclient.pool}, args{keyUnexistant}, TypeUint16, shp, valuesUint8, true},
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
			if tt.wantErr == false {
				if gotDt != tt.wantDt {
					t.Errorf("TensorGetValues() gotDt = %v, want %v", gotDt, tt.wantDt)
				}
				if !reflect.DeepEqual(gotShape, tt.wantShape) {
					t.Errorf("TensorGetValues() gotShape = %v, want %v", gotShape, tt.wantShape)
				}
				if !reflect.DeepEqual(gotData, tt.wantData) {
					t.Errorf("TensorGetValues() gotData = %v, want %v", gotData, tt.wantData)
				}
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
	valuesUint32 := []uint32{1}
	valuesUint64 := []uint64{1}

	keyFloat32 := "test:TensorSet:TypeFloat32:1"
	keyFloat64 := "test:TensorSet:TypeFloat64:1"

	keyInt8 := "test:TensorSet:TypeInt8:1"
	keyInt16 := "test:TensorSet:TypeInt16:1"
	keyInt32 := "test:TensorSet:TypeInt32:1"
	keyInt64 := "test:TensorSet:TypeInt64:1"

	keyByte := "test:TensorSet:Type[]byte:1"
	keyUint8 := "test:TensorSet:TypeUint8:1"
	keyUint16 := "test:TensorSet:TypeUint16:1"
	keyUint32 := "test:TensorSet:TypeUint32:ExpectError:1"
	keyUint64 := "test:TensorSet:TypeUint64:ExpectError:1"

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
		{keyFloat32, fields{pclient.pool}, args{keyFloat32, TypeFloat, shp, valuesFloat32}, false},
		{keyFloat64, fields{pclient.pool}, args{keyFloat64, TypeFloat64, shp, valuesFloat64}, false},

		{keyInt8, fields{pclient.pool}, args{keyInt8, TypeInt8, shp, valuesInt8}, false},
		{keyInt16, fields{pclient.pool}, args{keyInt16, TypeInt16, shp, valuesInt16}, false},
		{keyInt32, fields{pclient.pool}, args{keyInt32, TypeInt32, shp, valuesInt32}, false},
		{keyInt64, fields{pclient.pool}, args{keyInt64, TypeInt64, shp, valuesInt64}, false},

		{keyUint8, fields{pclient.pool}, args{keyUint8, TypeUint8, shp, valuesUint8}, false},
		{keyUint16, fields{pclient.pool}, args{keyUint16, TypeUint16, shp, valuesUint16}, false},
		{keyUint32, fields{pclient.pool}, args{keyUint32, TypeUint8, shp, valuesUint32}, true},
		{keyUint64, fields{pclient.pool}, args{keyUint64, TypeUint8, shp, valuesUint64}, true},

		{keyInt8Meta, fields{pclient.pool}, args{keyInt8Meta, TypeUint8, shp, nil}, false},
		{keyByte, fields{pclient.pool}, args{keyByte, TypeUint8, shp, valuesByte}, false},

		{"test:TestClient_TensorSet:1:FaultyDims", fields{pclient.pool}, args{"test:TestClient_TensorSet:1:FaultyDims", TypeFloat, []int{1, 10}, []float32{1}}, true},
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
		name      string
		args      args
		wantC     *Client
		wantError bool
	}{
		//{"test:Connect:BadUrl:1", args{"badurl",nil}, nil, true },

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
