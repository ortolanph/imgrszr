package maskable

import (
	"image"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"os"
	"fmt"

	"github.com/nfnt/resize"
)

func Maskable(inputPath string, outputPath string, width uint, height uint) error {
    fmt.Println(inputPath)
    fmt.Println(outputPath)
    fmt.Println(width)
    fmt.Println(height)

	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	cornerRadius := int(width) / 8

	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)
	
	maskedImage := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	maskRect := image.Rect(
		cornerRadius, 
		cornerRadius, 
		int(width) - cornerRadius, 
		int(width) - cornerRadius)
	draw.DrawMask(
		maskedImage, 
		img.Bounds(), 
		resizedImg, 
		image.Point{}, 
		maskRect, 
		image.Point{}, 
		draw.Src)

	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, maskedImage, nil)
	if err != nil {
		panic(err)
	}


	return nil;
}