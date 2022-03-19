// Package analyse extracts textual clacks messages from gocv.Mat
package analyse

import (
	"gocv.io/x/gocv"
	"image/color"
	"math/big"
)

// Color of the marked points
var pointColor = color.RGBA{0, 255, 255, 0}

func GetBytes(mat *gocv.Mat) string {
	if mat.Empty() {
		return ""
	}
	thresh := gocv.NewMat()
	defer thresh.Close()

	var weakThreshold float32 = 127
	var strongThreshold float32 = 255
	gocv.Threshold(*mat, &thresh, weakThreshold, strongThreshold, gocv.ThresholdBinary)

	thresh.CopyTo(mat)

	size := mat.Size()
	pixelWidth := size[1] / 2
	pixelHeight := size[0] / 4
	pixelXOffset := pixelWidth / 2
	pixelYOffset := pixelHeight / 2

	ch := mat.Channels()
	v := make([]uint8, ch)

	y := pixelYOffset
	x := pixelXOffset

	bi := big.NewInt(0)

	var b uint = 0

	for row := 0; row < 4; row++ {
		for col := 0; col < 2; col++ {
			y = row*pixelHeight + pixelYOffset
			x = col*pixelWidth + pixelXOffset

			for c := 0; c < ch; c++ {
				//fmt.Printf("%v", c)
				v[c] = mat.GetUCharAt(y, x*ch+c)
			}

			if v[0] == 255 {
				b = 1
			} else {
				b = 0
			}

			bi.SetBit(bi, row*2+col, b)
		}
	}

	return string(bi.Bytes())
}
