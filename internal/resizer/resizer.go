package resizer

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// ResizeImage resizes an image to the specified width and height
func ResizeImage(inputPath string, outputPath string, width, height int) error {
	// Open the source image
	srcFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("error opening source image: %v", err)
	}
	defer srcFile.Close()

	// Decode the image
	src, _, err := image.Decode(srcFile)
	if err != nil {
		return fmt.Errorf("error decoding image: %v", err)
	}

	// Create a new resized image
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	// Scale and draw the source image to the destination
	draw.NearestNeighbor.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)

	// Create output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outFile.Close()

	// Encode and save the resized image
	ext := strings.ToLower(filepath.Ext(outputPath))
	switch ext {
	case ".png":
		err = png.Encode(outFile, dst)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(outFile, dst, nil)
	default:
		return fmt.Errorf("unsupported file format: %s", ext)
	}

	if err != nil {
		return fmt.Errorf("error encoding image: %v", err)
	}

	return nil
}