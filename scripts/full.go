package scripts

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"gocv.io/x/gocv"

	rai "github.com/itamarhaber/redisai-go/redisai"
)

var cocoLabels = []string{
	"person",
	"bicycle",
	"car",
	"motorbike",
	"aeroplane",
	"bus",
	"train",
	"truck",
	"boat",
	"traffic light",
	"fire hydrant",
	"stop sign",
	"parking meter",
	"bench",
	"bird",
	"cat",
	"dog",
	"horse",
	"sheep",
	"cow",
	"elephant",
	"bear",
	"zebra",
	"giraffe",
	"backpack",
	"umbrella",
	"handbag",
	"tie",
	"suitcase",
	"frisbee",
	"skis",
	"snowboard",
	"sports ball",
	"kite",
	"baseball bat",
	"baseball glove",
	"skateboard",
	"surfboard",
	"tennis racket",
	"bottle",
	"wine glass",
	"cup",
	"fork",
	"knife",
	"spoon",
	"bowl",
	"banana",
	"apple",
	"sandwich",
	"orange",
	"broccoli",
	"carrot",
	"hot dog",
	"pizza",
	"donut",
	"cake",
	"chair",
	"sofa",
	"pottedplant",
	"bed",
	"diningtable",
	"toilet",
	"tvmonitor",
	"laptop",
	"mouse",
	"remote",
	"keyboard",
	"cell phone",
	"microwave",
	"oven",
	"toaster",
	"sink",
	"refrigerator",
	"book",
	"clock",
	"vase",
	"scissors",
	"teddy bear",
	"hair drier",
	"toothbrush",
}

func processImage(img gocv.Mat, height int) (cimg gocv.Mat) {
	// Utility to resize a rectangular image to a padded square (letterbox) and normalize it
	color := color.RGBA{127, 127, 127, 0}
	sz := img.Size()
	ih, iw := float64(sz[0]), float64(sz[1])
	imax := math.Max(ih, iw)
	ratio := float64(height) / imax
	ch, cw := int(ih*ratio), int(iw*ratio)
	if math.Mod(float64(ch), 2) != 0 {
		ch++
	}
	if math.Mod(float64(cw), 2) != 0 {
		cw++
	}
	nsz := image.Point{cw, ch}
	dh, dw := float64((height-ch)/2.0), float64((height-cw)/2.0)
	top, bottom := int(math.Round(dh-0.1)), int(math.Round(dh+0.1))
	left, right := int(math.Round(dw-0.1)), int(math.Round(dw+0.1))
	cimg = gocv.NewMat()

	gocv.Resize(img, &cimg, nsz, 0, 0, gocv.InterpolationLinear)
	gocv.CopyMakeBorder(cimg, &cimg, top, bottom, left, right, gocv.BorderConstant, color)
	gocv.CvtColor(cimg, &cimg, gocv.ColorBGRToRGB)
	cimg.ConvertTo(&cimg, gocv.MatTypeCV32F)
	cimg.DivideFloat(256.0)
	return cimg
}

func main() {
	height := 416
	filename := "dog.jpg"
	prefix := "frame:1234"

	window := gocv.NewWindow("Hello")
	img := gocv.IMRead(filename, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Error reading image from: %s\n", filename)
		return
	}
	cimg := processImage(img, height)

	rc := rai.Connect("redis://localhost:6379")

	// keyRAIModel := "rai:model:yolo-full"
	// // Load model
	// if err := rc.ModelSetFromFile(keyRAIModel,
	// 	rai.BackendTF, rai.DeviceCPU,
	// 	"/home/itamar/work/EdgeRealtimeVideoAnalytics/app/models/yolo_full_with_pipeline.pb",
	// 	[]string{"input_1", "input_image_shape"},
	// 	[]string{"concat_13", "concat_12", "concat_11"}); err != nil {
	// 	fmt.Printf("Error setting model '%v'\n", err)
	// 	return
	// }

	keyRAIModel := "rai:model:yolo-tiny"
	// Load model
	// if err := rc.ModelSetFromFile(keyRAIModel,
	// 	rai.BackendTF, rai.DeviceCPU,
	// 	"./models/tiny-yolo-voc.pb",
	// 	[]string{"input"},
	// 	[]string{"output"}); err != nil {
	// 	fmt.Printf("Error setting model '%v'\n", err)
	// 	return
	// }

	keyRAIScript := "rai:script:yolo-tiny"
	// Load model
	if err := rc.ScriptSetFromFile(keyRAIScript,
		rai.DeviceCPU,
		"./scripts/yolo-boxes.py"); err != nil {
		fmt.Printf("Error setting model '%v'\n", err)
		return
	}

	// // Set model's input shape tensor
	// keyRAIModelShapeTensor := "rai:tensor:yolo-shape"
	// if err := rc.TensorSet(keyRAIModelShapeTensor,
	// 	rai.TypeFloat, []int{2}, []int{height, height}); err != nil {
	// 	fmt.Printf("Error setting shape tensor '%v'\n", err)
	// 	return
	// }

	// // Set model's input tensor from processed image
	// keyInputTensor := fmt.Sprintf("%s:tensor:input", prefix)
	// if err := rc.TensorSet(keyInputTensor,
	// 	rai.TypeFloat,
	// 	[]int{1, cimg.Cols(), cimg.Rows(), cimg.Channels()},
	// 	cimg.ToBytes()); err != nil {
	// 	fmt.Printf("Error setting input tensor '%v'\n", err)
	// 	return
	// }

	// Set model's input tensor from processed image
	keyModelInputTensor := fmt.Sprintf("%s:tensor:model_input", prefix)
	if err := rc.TensorSet(keyModelInputTensor,
		rai.TypeFloat,
		[]int{1, cimg.Cols(), cimg.Rows(), cimg.Channels()},
		cimg.ToBytes()); err != nil {
		fmt.Printf("Error setting input tensor '%v'\n", err)
		return
	}

	// // Run the model
	// keyOutputTensor := fmt.Sprintf("%s:tensor:output:classes", prefix)
	// keyOutputTensorThings := fmt.Sprintf("%s:tensor:output:things", prefix)
	// keyOutputTensorBoxes := fmt.Sprintf("%s:tensor:output:boxes", prefix)
	// if err := rc.ModelRun(keyRAIModel,
	// 	[]string{keyInputTensor, keyRAIModelShapeTensor},
	// 	[]string{keyOutputTensorClasses, keyOutputTensorThings, keyOutputTensorBoxes}); err != nil {
	// 	fmt.Printf("Error running model ''%v'\n", err)
	// 	return
	// }

	// Run the model
	keyModelOutputTensor := fmt.Sprintf("%s:tensor:model_output", prefix)
	if err := rc.ModelRun(keyRAIModel,
		[]string{keyModelInputTensor},
		[]string{keyModelOutputTensor}); err != nil {
		fmt.Printf("Error running model ''%v'\n", err)
		return
	}

	// Run the script
	keyScriptOutputTensor := fmt.Sprintf("%s:tensor:script_output", prefix)
	if err := rc.ScriptRun(keyRAIScript, "boxes_from_tf",
		[]string{keyModelOutputTensor},
		[]string{keyScriptOutputTensor}); err != nil {
		fmt.Printf("Error running script ''%v'\n", err)
		return
	}

	_, dims, data, err := rc.TensorGetValues(keyScriptOutputTensor)
	if err != nil {
		fmt.Printf("Error running script ''%v'\n", err)
		return
	}

	// Crudely parse the resulting "tensor"
	// dims[0] should be == 1, i.e one-sized batch
	// dims[1] is the number of detected "boxes"
	// dims[2] is the number of elements of each "box"
	counts := make(map[int]int)
	for i := 0; i < dims[1]; i++ {
		box := data[i*dims[2] : (i+1)*dims[2]]
		if box[4] == 0.0 {
			// Ignore zero-confidence boxe
			continue
		}
		label := int(box[dims[2]-1])
		counts[label]++
	}
	fmt.Printf("%v\n", counts)

	// t := tensor.New(tensor.WithShape(ddims...), tensor.WithBacking(dvals))
	// s, _ := t.Slice(tensor.singleSlice(1))
	// for box := range t.Tensor3() {
	// 	fmt.Printf("%v\n", box.Shape())
	// }

	for {
		window.IMShow(cimg)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
