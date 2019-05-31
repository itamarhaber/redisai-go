package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func processImage(img gocv.Mat) (cimg gocv.Mat) {
	cimg = img.Clone()
	gocv.CvtColor(cimg, cimg, gocv.ColorBGRToGray)
	return cimg
}

func main() {
	filename := "sample_dog_416.jpg"

	window := gocv.NewWindow("Hello")
	img := gocv.IMRead(filename, gocv.IMReadColor)
	if img.Empty() {
		fmt.Println("Error reading image from: %v", filename)
		return
	}
	cimg := processImage(img)
	for {
		window.IMShow(cimg)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
