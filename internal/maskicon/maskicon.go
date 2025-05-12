package maskicon

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

// CreateMaskableIcon creates a maskable icon from the input image
func CreateMaskableIcon(inputPath string, outputPath string, size int) error {
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

	// Create a new image with a transparent background
	dst := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.Draw(dst, dst.Bounds(), &image.Uniform{color.Transparent}, image.Point{}, draw.Src)

	// Calculate scaling to fit within the icon while maintaining aspect ratio
	srcBounds := src.Bounds()
	srcWidth := srcBounds.Dx()
	srcHeight := srcBounds.Dy()

	// Calculate scale to fit
	scale := float64(size) / float64(max(srcWidth, srcHeight))
	newWidth := int(float64(srcWidth) * scale)
	newHeight := int(float64(srcHeight) * scale)

	// Create a scaled version of the source image
	scaledImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.NearestNeighbor.Scale(scaledImg, scaledImg.Bounds(), src, src.Bounds(), draw.Over, nil)

	// Calculate position to center the scaled image
	x := (size - newWidth) / 2
	y := (size - newHeight) / 2

	// Draw the scaled image onto the destination
	draw.Draw(dst, image.Rect(x, y, x+newWidth, y+newHeight), scaledImg, image.Point{}, draw.Over)

	// Create output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outFile.Close()

	// Encode as PNG (preferred for maskable icons)
	err = png.Encode(outFile, dst)
	if err != nil {
		return fmt.Errorf("error encoding image: %v", err)
	}

	return nil
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}