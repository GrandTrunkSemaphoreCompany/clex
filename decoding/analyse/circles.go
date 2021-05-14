package analyse

import (
	"github.com/GrandTrunkSemaphoreCompany/clex/decoding/overlay"
	"gocv.io/x/gocv"
	"image"
	"math"
)

var dp = 1.0
var cannyUpperThreshold = 75.0
var minRadius = 10
var maxRadius = 20
var accumulator = 20.0

func HoughCircles(img *gocv.Mat) {

	gray := gocv.NewMat()
	defer gray.Close()

	gocv.CvtColor(*img, &gray, gocv.ColorBGRToGray)


	circles := gocv.NewMat()
	defer circles.Close()

	gocv.HoughCirclesWithParams(gray, &circles, gocv.HoughGradient,  dp, float64(gray.Rows()/8), cannyUpperThreshold, accumulator, minRadius, maxRadius)


	//overlay.ApplyCircles(img, circles)

	bbox := CirclesToBoxelBoundingBox(circles)
	//overlay.CropBoundingBox(img, bbox)
	overlay.ApplyBoundingBox(img, bbox)
}

func CirclesToBoxelBoundingBox(circles gocv.Mat) image.Rectangle {

	if circles.Cols() < 4 {
		return image.Rectangle{Min: image.Point{X:0, Y:0}, Max: image.Point{X:0, Y:0}}
	}

	//var vecf []gocv.Vecf
	//

	var xMin float32 = float32(math.Inf(1))
	var xMax float32
	var yMin float32 = float32(math.Inf(1))
	var yMax float32
	var rAvg float32
	//
	for i := 0; i< circles.Cols(); i++ {
		//circles.GetVec
		p := circles.GetVecfAt(0, i)

		var px float32 = p[0]
		var py float32 = p[1]
		var pr float32 = p[2]

		if math.Round(float64(px/10)) < math.Round(float64(xMin/10)) {
			xMin = px
		}
		if math.Round(float64(px/10)) > math.Round(float64(xMax/10)) {
			xMax = px
		}
		if math.Round(float64(py/10)) < math.Round(float64(yMin/10)) {
			yMin = py
		}
		if math.Round(float64(py/10)) > math.Round(float64(yMax/10)) {
			yMax = py
		}

		rAvg += pr
	}

	rAvg = rAvg / float32(circles.Cols())

	//fmt.Println(xMin)
	//fmt.Println(xMax)
	//fmt.Println(yMin)
	//fmt.Println(yMax)
	//
	return image.Rectangle{Min: image.Point{X:int(xMin+2*rAvg), Y:int(yMin+2*rAvg)}, Max: image.Point{X:int(xMax-2*rAvg), Y:int(yMax-2*rAvg)}}
	////
	////fmt.Println(circles.GetVecfAt(0, 0))
	////fmt.Println(circles.GetVecdAt(0, 0))
	////fmt.Println(circles.GetVeciAt(0, 0))
	////fmt.Println(circles.GetVecbAt(0, 0))
	////fmt.Println(circles.GetVecfAt(0, 1))
	////fmt.Println(circles.GetVecfAt(0, 2))
	////fmt.Println(circles.GetVecfAt(0, 3 ))
	//
	//vecf := []gocv.Vecf{circles.GetVecfAt(0, 0),circles.GetVecfAt(0, 1),circles.GetVecfAt(0, 2),circles.GetVecfAt(0, 3)}
	//fmt.Printf("%v\n", vecf)
	//
	//sort.Slice(vecf, func(i, j int) bool {
	//	var x1 float64 = math.Round(float64(vecf[i][0]/10))
	//	var y1 float64 = math.Round(float64(vecf[i][1]/10))
	//	var x2 float64 = math.Round(float64(vecf[j][0]/10))
	//	var y2 float64 = math.Round(float64(vecf[j][1]/10))
	//
	//	fmt.Printf("%f %f - %f %f\n", x1, y1, x2, y2)
	//	fmt.Printf("%t - %f / %f\n", (x1 <= x2), x1, x2)
	//	fmt.Printf("%t - %f / %f\n", (y1 <= y2), y1, y2)
	//	return (x1 <= x2) && (y1 <= y2)
	//})
	//fmt.Printf("%v\n", vecf)
	//
	////xMin := math.Minvecf[0][0]
	////xMax := vecf[1][0]
	////yMin := vecf[0][0]
	////yMax := vecf[1][0]
	//
	//return image.Rectangle{Min: image.Point{X:0, Y:0}, Max: image.Point{X:0, Y:0}}


}