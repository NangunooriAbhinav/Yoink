package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func loadImage(filePath string) (image.Image, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return nil, "", err
	}

	return img, format, nil
}

func convert(img image.Image, format string, outputFilePath string) error {
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	switch format {
	case "jpeg":
		options := jpeg.Options{Quality: 100}
		return jpeg.Encode(outputFile, img, &options)
	case "png":
		return png.Encode(outputFile, img)
	default:
		return errors.New("unsupported format")
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go [input file path] [output file path] [-j/-p]")
		os.Exit(1)
	}

	filePath := os.Args[1]
	outputFilePath := os.Args[2]
	var outputFormat string

	if len(os.Args) == 4 {
		switch os.Args[3] {
		case "-j":
			outputFormat = "jpeg"
		case "-p":
			outputFormat = "png"
		default:
			fmt.Println("Invalid format flag. Use -j for JPEG or -p for PNG.")
			os.Exit(1)
		}
	} else {
		fmt.Println("Please specify the output format flag: -j for JPEG or -p for PNG.")
		os.Exit(1)
	}

	if outputFormat == "" {
		fmt.Println("No output format specified.")
		os.Exit(1)
	}

	img, _, err := loadImage(filePath)
	if err != nil {
		fmt.Println("Error loading image:", err)
		os.Exit(1)
	}

	ext := filepath.Ext(outputFilePath)
	if ext != "."+outputFormat {
		outputFilePath = strings.TrimSuffix(outputFilePath, ext) + "." + outputFormat
	}

	err = convert(img, outputFormat, outputFilePath)
	if err != nil {
		fmt.Println("Error converting image:", err)
		os.Exit(1)
	}

	fmt.Printf("Image successfully converted and saved as %s\n", outputFilePath)
}
