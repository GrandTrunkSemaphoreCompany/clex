package analyse

import (
	"fmt"
	"github.com/GrandTrunkSemaphoreCompany/clex/pkg/decoding/overlay"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"math"
)

var lineColor = color.RGBA{255, 0, 0, 0}

func FindAreasAndFormat(img *gocv.Mat, analyzedImg *gocv.Mat) {
	blurImg := gocv.NewMat()
	defer blurImg.Close()

	greyImg := gocv.NewMat()
	defer greyImg.Close()

	thresholdImg := gocv.NewMat()
	defer thresholdImg.Close()

	gocv.CvtColor(*img, &greyImg, gocv.ColorBGRToGray)
	erodeKernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3))
	gocv.Erode(greyImg, &greyImg, erodeKernel)
	gocv.Erode(greyImg, &greyImg, erodeKernel)
	gocv.Blur(greyImg, &blurImg, image.Point{5, 5})
	gocv.Threshold(blurImg, &thresholdImg, 0.0, 255.0, gocv.ThresholdBinary+gocv.ThresholdOtsu)

	// Step 2 Connect Individual Contours

	//kernel = cv2.getStructuringElement(cv2.MORPH_RECT, (5,5))
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Point{5, 5})
	defer kernel.Close()

	//close = cv2.morphologyEx(thresh, cv2.MORPH_CLOSE, kernel, iterations=2)
	gocv.MorphologyEx(thresholdImg, analyzedImg, gocv.MorphClose, kernel)

	//Step 3 Filter for Code
	contours := gocv.FindContours(thresholdImg, gocv.RetrievalExternal, gocv.ChainApproxSimple)

	overlay.ApplyTracing(img, contours)

	matCanny := gocv.NewMat()
	matLines := gocv.NewMat()

	gocv.Canny(thresholdImg, &matCanny, 100, 100)
	gocv.HoughLinesP(matCanny, &matLines, 0.5, math.Pi/360, 40)

	fmt.Println(matLines.Cols())
	fmt.Println(matLines.Rows())

	if !matLines.Empty() {
		for index1 := 0; index1 < matLines.Rows(); index1++ {
			pt1 := image.Pt(int(matLines.GetVeciAt(index1, 0)[0]), int(matLines.GetVeciAt(index1, 0)[1]))
			pt2 := image.Pt(int(matLines.GetVeciAt(index1, 0)[2]), int(matLines.GetVeciAt(index1, 0)[3]))
			gocv.Line(analyzedImg, pt1, pt2, lineColor, 10)
		}
	} else {
		fmt.Println("No lines found")
	}
}

//Convert image to grayscale and median blur to smooth image
//Sharpen image to enhance edges
//Threshold
//Perform morphological transformations
//Find contours and filter using minimum/maximum threshold area
//Crop and save ROI
func FindContoursAndROI(img *gocv.Mat) {
	greyImg := gocv.NewMat()
	defer greyImg.Close()

	blurImg := gocv.NewMat()
	defer blurImg.Close()

	thresholdImg := gocv.NewMat()
	defer thresholdImg.Close()

	gocv.CvtColor(*img, &greyImg, gocv.ColorBGRToGray)
	gocv.Blur(greyImg, &blurImg, image.Point{5, 5})

	//sharpImg := blurImg
	//sharpen_kernel = np.array([[-1,-1,-1], [-1,9,-1], [-1,-1,-1]])
	//sharpen = cv2.filter2D(blur, -1, sharpen_kernel)

	//	thresh = cv2.threshold(sharpen,160,255, cv2.THRESH_BINARY_INV)[1]
	gocv.Threshold(blurImg, img, 0.0, 255.0, gocv.ThresholdBinary+gocv.ThresholdOtsu)

	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Point{5, 5})
	defer kernel.Close()

	//gocv.Erode(thresholdImg, &thresholdImg, kernel)

	//contours := gocv.FindContours(thresholdImg, gocv.RetrievalExternal, gocv.ChainApproxSimple )
	//overlay.ApplyTracing(img, contours)

	gocv.HoughLinesPWithParams(thresholdImg, img, math.Pi/180, 1, 1, 0, 0)

	//gocv.Dilate(*img, img, kernel)

	//*img = sharpImg

	/*
		kernel = cv2.getStructuringElement(cv2.MORPH_RECT, (3,3))
		close = cv2.morphologyEx(thresh, cv2.MORPH_CLOSE, kernel, iterations=2)

		cnts = cv2.findContours(close, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)
		cnts = cnts[0] if len(cnts) == 2 else cnts[1]

		min_area = 100
		max_area = 1500
		image_number = 0
		for c in cnts:
		    area = cv2.contourArea(c)
		    if area > min_area and area < max_area:
		        x,y,w,h = cv2.boundingRect(c)
		        ROI = image[y:y+h, x:x+h]
		        cv2.imwrite('ROI_{}.png'.format(image_number), ROI)
		        cv2.rectangle(image, (x, y), (x + w, y + h), (36,255,12), 2)
		        image_number += 1

		cv2.imshow('sharpen', sharpen)
		cv2.imshow('close', close)
		cv2.imshow('thresh', thresh)
		cv2.imshow('image', image)
		cv2.waitKey()

	*/

}

func FindByChessboard(img *gocv.Mat) {

	cornersMat := gocv.NewMat()
	defer cornersMat.Close()

	found := gocv.FindChessboardCorners(*img, image.Point{9, 6}, &cornersMat, gocv.CalibCBAdaptiveThresh&gocv.CalibCBNormalizeImage&gocv.CalibCBExhaustive)
	fmt.Println(found)
	gocv.DrawChessboardCorners(img, image.Point{9, 6}, cornersMat, true)

	//gocv.
	//found := gocv.FindChessboardCorners(*img, image.Point{9, 6}, &cornersMat, gocv.CalibCBAdaptiveThresh & gocv.CalibCBNormalizeImage & gocv.CalibCBExhaustive)
	//fmt.Println(found)
	//gocv.DrawChessboardCorners(img, image.Point{9,6}, cornersMat, true)

}

func Manipulate(img *gocv.Mat) {
	//img := gocv.NewMat()
	//img = src.Clone()

	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(*img, &gray, gocv.ColorBGRToGray)

	edges := gocv.NewMat()
	defer edges.Close()

	var weakThreshold float32 = 300
	var strongThreshold float32 = 400
	gocv.Canny(gray, &edges, weakThreshold, strongThreshold)

	//Step 3 Filter for Code
	contours := gocv.FindContours(edges, gocv.RetrievalExternal, gocv.ChainApproxSimple)

	overlay.ApplyTracing(img, contours)

}

func HoughTransform(img *gocv.Mat) {
	//img := gocv.NewMat()
	//img = src.Clone()

	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(*img, &gray, gocv.ColorBGRToGray)

	edges := gocv.NewMat()
	defer edges.Close()

	var weakThreshold float32 = 300
	var strongThreshold float32 = 400
	gocv.Canny(gray, img, weakThreshold, strongThreshold)

	lines := gocv.NewMat()
	defer lines.Close()

	var rho float32 = 1
	var theta float32 = math.Pi / 180
	var threshold int = 100
	var minLineLength float32 = 90
	var maxLineGap float32 = 100
	if !edges.Empty() {
		gocv.HoughLinesPWithParams(edges, &lines, rho, theta, threshold, minLineLength, maxLineGap)
	}

	if !lines.Empty() {
		fmt.Printf("Line count: %d\n", lines.Rows())
		for idx := 0; idx < lines.Rows(); idx++ {
			line := lines.GetVeciAt(idx, 0)
			gocv.Line(img, image.Point{int(line[0]), int(line[1])}, image.Point{int(line[2]), int(line[3])}, color.RGBA{0, 255, 0, 0}, 2)
		}
	} else {
		//if there are no lines, return Empty Mat
		fmt.Println("No lines found")
		//img = gocv.NewMat()
	}

	if lines.Rows() == 16 {
		fmt.Printf("Found all lines\n")
		line := lines.GetVeciAt(0, 0)
		fmt.Println(line[0], line[1], line[2], line[3])
		line = lines.GetVeciAt(1, 0)
		fmt.Println(line[0], line[1], line[2], line[3])
		line = lines.GetVeciAt(1, 1)
		fmt.Println(line[0], line[1], line[2], line[3])
	}

	//return img
}
