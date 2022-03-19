package cmd

import (
	"fmt"
	"github.com/GrandTrunkSemaphoreCompany/clex/decoding/analyse"
	"github.com/GrandTrunkSemaphoreCompany/clex/decoding/display"
	"github.com/GrandTrunkSemaphoreCompany/clex/decoding/overlay"
	"gocv.io/x/gocv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(decodeCmd)
}

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decodes a Clacks message",
	Run: func(cmd *cobra.Command, args []string) {
		var captures = []VideoInput{}

		pos := 1
		for _, c := range cameras {
			capture, err := c.InitVideoInput()

			if err != nil {
				fmt.Printf("Error configuring video capture device: %s (id %d)\n", c.Url, c.Id)
				continue
			}
			capture.CameraConfig.Position = pos
			captures = append(captures, capture)
			pos++
		}

		fmt.Println("------")
		for _, c := range captures {
			fmt.Printf("Camera: %s\n", c.CameraConfig.Url)
			fmt.Printf("\t%f x %f\n", c.VideoCapture.Get(gocv.VideoCaptureFrameWidth), c.VideoCapture.Get(gocv.VideoCaptureFrameHeight))
		}

		fmt.Println("------")

		window := gocv.NewWindow("Display")
		defer window.Close()

		scanCameras(captures, window)
		fmt.Println("decode: not implemented")
	},
}

func scanCameras(cameras []VideoInput, window *gocv.Window) {
	windowMat := gocv.NewMatWithSizeFromScalar(gocv.Scalar{127, 0, 0, 0}, 480, 640, gocv.MatTypeCV8UC3)
	defer windowMat.Close()

	workingMat := gocv.NewMat()
	defer workingMat.Close()

	for {
		//fmt.Println("> Running loop")

		for _, c := range cameras {
			if ok := c.VideoCapture.Read(&workingMat); !ok {
				fmt.Printf("Device closed: %v\n", c.CameraConfig.Url)
				fmt.Println("")
				continue
			}

			if workingMat.Empty() {
				continue
			}

			display.DisplayInWindow(window, &windowMat, c.CameraConfig.Position, 1, len(cameras), &workingMat)

			analyse.HoughCircles(&workingMat)
			text := analyse.GetBytes(&workingMat)
			overlay.ApplyText(&workingMat, text)
			overlay.ApplyFPS(&workingMat)

			display.DisplayInWindow(window, &windowMat, c.CameraConfig.Position, 2, len(cameras), &workingMat)
		}
	}
}
