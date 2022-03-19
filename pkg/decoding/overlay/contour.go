package overlay

import (
	//"fmt"
	"gocv.io/x/gocv"
)

const MinimumArea = 1

// Applies a countour line around supplied contours
func ApplyTracing(img *gocv.Mat, contours gocv.PointsVector) {
	//fmt.prin

	for i := 0; i < contours.Size(); i++ {
		c := contours.At(i)
		//fmt.Println(i)
		//fmt.Println(c)
		area := gocv.ContourArea(c)
		//area := gocv.ContourArea(c)
		if area < MinimumArea {
			continue
		}

		//status: = "Edges detected"
		gocv.DrawContours(img, contours, i, contourColor, 2)
		rect := gocv.BoundingRect(c)
		gocv.Rectangle(img, rect, bboxColor, 2)

		rotrect := gocv.MinAreaRect(c)
		vertices := rotrect.Points
		for i = 0; i < 4; i++ {
			gocv.Line(img, vertices[i], vertices[(i+1)%4], rotatedbboxColor, 2)
		}

		//gocv.BoundingRect(contours)
		//gocv.Rectangle(img, rect, bboxColor, 2)

	}
}
