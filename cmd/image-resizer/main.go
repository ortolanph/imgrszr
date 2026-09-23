package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ortolanph/imgrszr/internal/maskable"
	"github.com/ortolanph/imgrszr/internal/resizer"
)

func main() {
	// Resize command flags
	resizeCmd := flag.NewFlagSet("resize", flag.ExitOnError)
	inputPath := resizeCmd.String("input", "", "Path to the input image")
	outputPath := resizeCmd.String("output", "", "Path to the output image")
	width := resizeCmd.Uint("width", 0, "Desired width of the image")
	height := resizeCmd.Uint("height", 0, "Desired height of the image")

	maskCmd := flag.NewFlagSet("mask", flag.ExitOnError)
	maskInputPath := maskCmd.String("input", "", "Path to the input image")
	maskOutputPath := maskCmd.String("output", "", "Path to the output image")
	maskWidth := maskCmd.Uint("width", 0, "Desired width of the image")
	maskHeight := maskCmd.Uint("height", 0, "Desired height of the image")

	// Check if a subcommand was provided
	if len(os.Args) < 2 {
		fmt.Println("Expected 'resize' or 'mask' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "resize":
		resizeCmd.Parse(os.Args[2:])

		// Validate resize inputs
		if *inputPath == "" || *outputPath == "" || *width <= 0 || *height <= 0 {
			fmt.Println("Usage: imgrszr resize -input <input-file> -output <output-file> -width <width> -height <height>")
			os.Exit(1)
		}

		err := resizer.ResizeImage(*inputPath, *outputPath, *width, *height)
		if err != nil {
			log.Fatalf("Error resizing image: %v", err)
		}
		fmt.Printf("Image resized and saved to %s\n", *outputPath)

	case "mask":
		maskCmd.Parse(os.Args[2:])

		// Validate resize inputs
		if *maskInputPath == "" || *maskOutputPath == "" || *maskWidth <= 0 || *maskHeight <= 0 {
			fmt.Println("Usage: imgrszr mask -input <input-file> -output <output-file> -width <width> -height <height>")
			os.Exit(1)
		}

		err := maskable.Maskable(*maskInputPath, *maskOutputPath, *maskWidth, *maskHeight)
		if err != nil {
			log.Fatalf("Error resizing image: %v", err)
		}
		fmt.Printf("Image resized and saved to %s\n", *maskOutputPath)

	default:
		fmt.Println("Expected 'resize' or 'mask' subcommands")
		os.Exit(1)
	}
}
