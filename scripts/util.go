package scripts

// func processImage(img gocv.Mat, height int) (cimg gocv.Mat) {
// 	// Utility to resize a rectangular image to a padded square (letterbox) and normalize it
// 	color := color.RGBA{127, 127, 127, 0}
// 	sz := img.Size()
// 	ih, iw := float64(sz[0]), float64(sz[1])
// 	imax := math.Max(ih, iw)
// 	ratio := float64(height) / imax
// 	ch, cw := int(math.Round(ih*ratio)), int(math.Round(iw*ratio))
// 	nsz := image.Point{cw, ch}
// 	dh, dw := float64((height-ch)/2.0), float64((height-cw)/2.0)
// 	top, bottom := int(math.Round(dh-0.1)), int(math.Round(dh+0.1))
// 	left, right := int(math.Round(dw-0.1)), int(math.Round(dw+0.1))
// 	cimg = gocv.NewMat()

// 	gocv.Resize(img, &cimg, nsz, 0, 0, gocv.InterpolationLinear)
// 	gocv.CopyMakeBorder(cimg, &cimg, top, bottom, left, right, gocv.BorderConstant, color)
// 	cimg.ConvertTo(&cimg, gocv.MatTypeCV32F)
// 	cimg.DivideFloat(255.0)
// 	return cimg
// }
