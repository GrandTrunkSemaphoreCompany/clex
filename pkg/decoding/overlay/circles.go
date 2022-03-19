// Package overlay applies visual identifiers over an gocv.Mat to help with
// visual debugging of analysis code
package overlay

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
)

// Adds a bounding box to a gocv.Mat using bbox
func ApplyBoundingBox(img *gocv.Mat, bbox image.Rectangle) {
	gocv.Rectangle(img, bbox, bboxColor, 2)

}

func CropBoundingBox(img *gocv.Mat, bbox image.Rectangle) {
	cropped := img.Region(bbox)
	cropped.CopyTo(img)
	//gocv.Rectangle(img, bbox, bboxColor, 2)

}

// Adds sized circles to a gocv.Mat using circles
func ApplyCircles(img *gocv.Mat, circles gocv.Mat) {
	if circles.Empty() {
		fmt.Errorf("Empty HoughCircles test")
	}
	if circles.Rows() != 1 {
		fmt.Errorf("Invalid HoughCircles test rows: %v", circles.Rows())
	}
	if circles.Cols() < 317 || circles.Cols() > 334 {
		fmt.Errorf("Invalid HoughCircles test cols: %v", circles.Cols())
	}

	fmt.Printf("Rows: %v / Cols: %v\n", circles.Rows(), circles.Cols())

	//for i: = 0; i < circles.Size(); i++ {
	//for i, _ := range circles.Size() {
	for i := 0; i < circles.Cols(); i++ {
		v := circles.GetVecfAt(0, i)
		if len(v) > 2 {
			x := int(v[0])
			y := int(v[1])
			r := int(v[2])

			fmt.Printf("%d = %d @ %d / %d\n", i, r, x, y)
			gocv.Circle(img, image.Pt(x, y), r, contourColor, 2)
			gocv.Circle(img, image.Pt(x, y), 2, bboxColor, 3)
		}
		//circles.
		//var x int32 = circles.GetIntAt(i, 0)
		//var y int32 =  circles.GetIntAt(i, 1)
		//var radius int32 = circles.GetIntAt(i, 2)
		//
		//if(x <=0 || y <=0 || radius <=0) {
		//	continue
		//}
		//fmt.Printf("(%d) %d - %d @ %d\n", i, x, y, radius)
		//var center = image.Point{int(x), int(y)}
		//
		//// draw the circle center
		//gocv.Circle(img, center, 3, contourColor, 1)
		//// draw the circle outline
		//gocv.Circle(img, center, int(radius), contourColor, 2)
	}
}
