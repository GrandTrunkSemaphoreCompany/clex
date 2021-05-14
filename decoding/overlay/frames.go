package overlay

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"time"
)

var measurement float64 = 0.0
var smoothing float64 = 0.5
var tickTime time.Time = time.Now()


// Adds a value smoothed frames per second count to the gocv.Mat
func ApplyFPS(baseImg *gocv.Mat ) {
	currentTime := time.Now()

	//fmt.Println(tickTime)
	//fmt.Println(currentTime)

	fpsDuration := currentTime.Sub(tickTime)
	//fmt.Printf("duration: %f\n", fpsDuration.Seconds())
	//fmt.Printf("duration: %d\n", fpsDuration.Milliseconds())

	measurement = (measurement * smoothing) + fpsDuration.Seconds() * (1.0 - smoothing)
	//fmt.Printf("measuement: %f\n", measurement)

	fpsLabel := fmt.Sprintf("%.0ffps", 1/measurement)

	gocv.PutText(baseImg, fpsLabel, image.Pt(10, 30), gocv.FontHersheySimplex, 1, fpsColor, 2)

	tickTime = currentTime
}
