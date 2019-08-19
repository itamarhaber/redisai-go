package scripts

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	// Partialy-meme-ified tiny YOLO vocabulary
	labels = []string{
		"MetalBirb",          //  0
		"Bicycle",            //  1
		"Birb",               //  2
		"Boat",               //  3
		"Bottle",             //  4
		"Bus",                //  5
		"Car",                //  6
		"Kitter",             //  7
		"Chair",              //  8
		"MilkPanda",          //  9
		"DiningTable",        // 10
		"Doggo",              // 11
		"ClipClopNeighDoggo", // 12
		"Motorbike",          // 13
		"Hooman",             // 14
		"PottedPlant",        // 15
		"PrairieCloud",       // 16
		"Sofa",               // 17
		"Train",              // 18
		"TVMonitor",          // 19
	}
)

// Client is the RedisAI client
type Client struct {
	pool *redis.Pool
}

// Initialize initializes
func Initialize() (c *Client, err error) {
	c = &Client{
		pool: &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial:        func() (redis.Conn, error) { return redis.DialURL("redis://localhost:6379") },
		},
	}

	return c, nil
}

// Check checks
func (c *Client) Check() error {
	conn := c.pool.Get()
	defer conn.Close()

	rep, err := conn.Do("PING")
	if err != nil {
		return err
	}
	fmt.Println(rep)
	return nil
}

// TensorSetFromBlob sets a tensor from a blob
func (c *Client) TensorSetFromBlob(key string, dt string, dims []int, blob []byte) (err error) {
	conn := c.pool.Get()
	defer conn.Close()

	args := redis.Args{}.Add(key, dt).AddFlat(dims).Add("BLOB", string(blob))
	rep, err := redis.String(conn.Do("AI.TENSORSET", args...))
	if err != nil {
		return err
	}
	if rep != "OK" {
		return fmt.Errorf("AI.TENSORSET failed - %s", rep)
	}
	return nil
}

// TensorGetAsValues gets a tensor as values
func (c *Client) TensorGetAsValues(key string) (dt string, dims []int, values []float64, err error) {
	conn := c.pool.Get()
	defer conn.Close()

	args := redis.Args{}.Add(key, "VALUES")
	rep, err := redis.MultiBulk(conn.Do("AI.TENSORGET", args...))
	if err != nil {
		return
	}
	dt, err = redis.String(rep[0], err)
	if err != nil {
		return
	}
	dims, err = redis.Ints(rep[1], err)
	if err != nil {
		return
	}
	values, err = redis.Float64s(rep[2], err)
	if err != nil {
		return
	}
	return
}

// ModelRun runs a model
func (c *Client) ModelRun(key string, inputs []string, outputs []string) (err error) {
	conn := c.pool.Get()
	defer conn.Close()

	args := redis.Args{}.Add(key, "INPUTS").AddFlat(inputs).Add("OUTPUTS").AddFlat(outputs)
	rep, err := redis.String(conn.Do("AI.MODELRUN", args...))
	if err != nil {
		return err
	}
	if rep != "OK" {
		return fmt.Errorf("AI.MODELRUN failed - %s", rep)
	}
	return nil
}

// ScriptRun runs a script
func (c *Client) ScriptRun(key string, fn string, inputs []string, outputs []string) (err error) {
	conn := c.pool.Get()
	defer conn.Close()

	args := redis.Args{}.Add(key, fn, "INPUTS").AddFlat(inputs).Add("OUTPUTS").AddFlat(outputs)
	rep, err := redis.String(conn.Do("AI.SCRIPTRUN", args...))
	if err != nil {
		return err
	}
	if rep != "OK" {
		return fmt.Errorf("AI.SCRIPTRUN failed - %s", rep)
	}
	return nil
}

// func main() {
// 	height := 416
// 	filename := "dogcat.jpg"

// 	window := gocv.NewWindow("Hello")
// 	img := gocv.IMRead(filename, gocv.IMReadColor)
// 	if img.Empty() {
// 		fmt.Printf("Error reading image from: %s\n", filename)
// 		return
// 	}
// 	cimg := processImage(img, height)
// 	rai, _ := Initialize()
// 	dims := []int{1, cimg.Cols(), cimg.Rows(), cimg.Channels()}
// 	err := rai.TensorSetFromBlob("in", "FLOAT", dims, cimg.ToBytes())
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	err = rai.ModelRun("yolo:model", []string{"in"}, []string{"out"})
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	err = rai.ScriptRun("yolo:script", "boxes_from_tf", []string{"out"}, []string{"post"})
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	_, ddims, dvals, err := rai.TensorGetAsValues("post")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// Crudely parse the resulting "tensor"
// 	// ddims[0] should be == 1, i.e one-sized batch
// 	// ddims[1] is the number of detected "boxes"
// 	// ddims[2] is the number of elements of each "box"
// 	counts := make(map[string]int)
// 	for i := 0; i < ddims[1]; i++ {
// 		box := dvals[i*ddims[2] : (i+1)*ddims[2]]
// 		if box[4] == 0.0 {
// 			// Ignore zero-confidence boxe
// 			continue
// 		}
// 		label := int(box[ddims[2]-1])
// 		counts[labels[label]]++
// 	}
// 	fmt.Printf("%v\n", counts)

// 	// t := tensor.New(tensor.WithShape(ddims...), tensor.WithBacking(dvals))
// 	// s, _ := t.Slice(tensor.singleSlice(1))
// 	// for box := range t.Tensor3() {
// 	// 	fmt.Printf("%v\n", box.Shape())
// 	// }

// 	for {
// 		window.IMShow(cimg)
// 		if window.WaitKey(1) >= 0 {
// 			break
// 		}
// 	}
// }
