package resizer

import (
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"os"

	"github.com/nfnt/resize"
)

// ResizeImage resizes an image to the specified width and height
func ResizeImage(inputPath string, outputPath string, width uint, height uint) error {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	resizedImg := resize.Resize(width, height, img, resize.Lanczos3) // Lanczos3 is a good interpolation for upscaling

	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, resizedImg, nil)
	if err != nil {
		panic(err)
	}

	return nil
}
