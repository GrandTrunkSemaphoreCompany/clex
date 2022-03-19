// Package display is used to display a number of gocv.Mat in a grid
package display

import (
	"gocv.io/x/gocv"
	"image"
	"os"
)

// Compose subset gocv.Mat into display gocv.Mat and show in window
func DisplayInWindow(window *gocv.Window, displayMat *gocv.Mat, row int, col int, maxRows int, subsetMat *gocv.Mat) {
	if !displayMat.Empty() && !subsetMat.Empty() {
		ComposeImage(row, col, maxRows, displayMat, subsetMat)
	}

	if !displayMat.Empty() {
		window.IMShow(*displayMat)
	}

	if window.WaitKey(1) == 27 {
		os.Exit(0)
	}
}

// Resize subset gocv.Mat into display gocv.Mat using rows and column positioning
func ComposeImage(row int, col int, maxRows int, displayMat *gocv.Mat, subsetMat *gocv.Mat) {
	var resizedWidth int = 640 / maxRows
	var resizedHeight int = 480 / 2

	coords := image.Rectangle{Min: image.Point{X: (row - 1) * resizedWidth, Y: (col - 1) * resizedHeight}, Max: image.Point{X: (row) * resizedWidth, Y: col * resizedHeight}}
	region := displayMat.Region(coords)

	gocv.Resize(*subsetMat, &region, coords.Size(), 0, 0, gocv.InterpolationLinear)
}
