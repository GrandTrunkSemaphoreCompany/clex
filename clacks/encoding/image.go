package encoding

import (
	"fmt"
	"github.com/disintegration/gift"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

type Image struct {
	Directory string
	Counter   int
}

func (ci *Image) New(d string) *Image {
	i := new(Image)
	i.Directory = d

	return i
}

func (ci *Image) Write(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		clexImage := ci.MakeClacksFromByte(p[i])

		filename := fmt.Sprintf("%s/%03d.png", ci.Directory, ci.Counter)

		myfile, err := os.Create(filename)
		if err != nil {
			return 0, err
		}

		png.Encode(myfile, clexImage)

		ci.Counter = ci.Counter + 1
	}

	return len(p), nil
}

func (ci *Image) AddClacksShutters(im *image.RGBA, bnd image.Rectangle, byte uint8) {
	spacer := 10
	var shutterWidth int
	var shutterHeight int

	ClacksWhite := color.RGBA{242, 243, 244, 255}
	ClacksBlack := color.RGBA{41, 36, 33, 255}

	shutterWidth = (bnd.Bounds().Dx() - (3 * spacer)) / 2
	shutterHeight = (bnd.Bounds().Dy() - (5 * spacer)) / 4

	var rects []image.Rectangle

	hPos := 0
	vPos := 0

	for h := 0; h < 4; h++ {
		vPos = bnd.Bounds().Min.Y + spacer + (h * shutterHeight) + (h * spacer)

		for w := 0; w < 2; w++ {
			hPos = bnd.Bounds().Min.X + spacer + (w * shutterWidth) + (w * spacer)

			rects = append(rects, image.Rect(hPos, vPos, hPos+shutterWidth, vPos+shutterHeight))
		}
	}

	// Start from least significant bit
	for i := 0; i < len(rects); i++ {
		var bitColor color.RGBA
		bit := byte & 1

		if bit == 1 {
			bitColor = ClacksWhite
		} else {
			bitColor = ClacksBlack
		}
		draw.Draw(im, rects[i], &image.Uniform{bitColor}, image.ZP, draw.Src)

		byte = byte >> 1
	}
}

func (ci *Image) AddSky(im *image.RGBA) {
	skyBlue := color.RGBA{135, 206, 235, 255}
	draw.Draw(im, im.Bounds(), &image.Uniform{skyBlue}, image.ZP, draw.Src)
}

func (ci *Image) AddClacksFrame(im *image.RGBA) image.Rectangle {
	skySpacer := 10
	ClacksBrown := color.RGBA{193, 154, 107, 255}
	ClacksRect := image.Rect(skySpacer, skySpacer, im.Bounds().Dx()-skySpacer, im.Bounds().Dy()-skySpacer) //  geometry of 2nd rectangle

	draw.Draw(im, ClacksRect, &image.Uniform{ClacksBrown}, image.ZP, draw.Src)

	return ClacksRect
}

func (ci *Image) MakeClacksFromByte(b uint8) image.Image {
	m := image.NewRGBA(image.Rect(0, 0, 130, 190)) // x1,y1,  x2,y2

	ci.AddSky(m)
	ClacksRect := ci.AddClacksFrame(m)
	ci.AddClacksShutters(m, ClacksRect, b)

	return m
}

func (ci *Image) Read(p []byte) (n int, err error) {
	// check directory set

	// Read file from directory
	// Decode

	// func (f *File) Readdir(n int) ([]FileInfo, error) {

	// load directory
	files, err := ioutil.ReadDir(ci.Directory)
	if err != nil {
		log.Fatal(err)
	}

	for i, f := range files {
		// load each file
		img, err := ReadFile(ci.Directory + "/" + f.Name())
		if err != nil {
			return i, err
		}

		b, err := GetByteFromImage(img)

		if err != nil {
			return i, err
		}

		p[i] = b
	}

	return len(files), nil
}

func ReadFile(filepath string) (image image.Image, error error) {
	// Read image from file that already exists
	existingImageFile, err := os.Open(filepath)
	if err != nil {
		// Handle error
	}
	defer existingImageFile.Close()

	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		// Handle error
	}

	return loadedImage, err
}

func GetByteFromImage(img image.Image) (byte byte, error error) {
	// 1. Create a new filter list and add some filters.
	g := gift.New(
		gift.Threshold(70),
	)

	dst := image.NewRGBA(g.Bounds(img.Bounds()))

	// 2. Use the Draw func to apply the filters to src and store the result in dst.
	g.Draw(dst, img)

	// 3. Sample each threshold pixel as a bbit
	t := make([]bool, 8)

	rr, _, _, _ := dst.At(40, 35).RGBA()
	t[7] = rr == 65535
	rr, _, _, _ = dst.At(90, 35).RGBA()
	t[6] = rr == 65535

	rr, _, _, _ = dst.At(40, 75).RGBA()
	t[5] = rr == 65535
	rr, _, _, _ = dst.At(90, 75).RGBA()
	t[4] = rr == 65535

	rr, _, _, _ = dst.At(40, 115).RGBA()
	t[3] = rr == 65535
	rr, _, _, _ = dst.At(90, 115).RGBA()
	t[2] = rr == 65535

	rr, _, _, _ = dst.At(40, 155).RGBA()
	t[1] = rr == 65535
	rr, _, _, _ = dst.At(90, 155).RGBA()
	t[0] = rr == 65535

	return BoolsToBytes(t)[0], nil
}
