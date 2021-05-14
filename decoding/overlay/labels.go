package overlay

import (
	"gocv.io/x/gocv"
	"image"
)

// Adds a status to the gocv.Mat
func ApplyText(baseImg *gocv.Mat, status string) {
	gocv.PutText(baseImg, status, image.Pt(120, 30), gocv.FontHersheyPlain, 2, textColor, 2)
}
