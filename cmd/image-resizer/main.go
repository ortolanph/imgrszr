package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ortolanph/imgrszr/internal/resizer"
)

func main() {
	// Resize command flags
	resizeCmd := flag.NewFlagSet("resize", flag.ExitOnError)
	inputPath := resizeCmd.String("input", "", "Path to the input image")
	outputPath := resizeCmd.String("output", "", "Path to the output image")
	width := resizeCmd.Uint("width", 0, "Desired width of the image")
	height := resizeCmd.Uint("height", 0, "Desired height of the image")

	// Check if a subcommand was provided
	if len(os.Args) < 2 {
		fmt.Println("Expected 'resize' or 'maskable' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "resize":
		resizeCmd.Parse(os.Args[2:])

		// Validate resize inputs
		if *inputPath == "" || *outputPath == "" || *width <= 0 || *height <= 0 {
			fmt.Println("Usage: program resize -input <input-file> -output <output-file> -width <width> -height <height>")
			os.Exit(1)
		}

		err := resizer.ResizeImage(*inputPath, *outputPath, *width, *height)
		if err != nil {
			log.Fatalf("Error resizing image: %v", err)
		}
		fmt.Printf("Image resized and saved to %s\n", *outputPath)

	default:
		fmt.Println("Expected 'resize' subcommands")
		os.Exit(1)
	}
}
