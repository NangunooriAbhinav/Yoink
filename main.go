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

func loadImage(filePath string) (image.Image,string,error){
	file, err := os.Open(filePath)

	if err != nil{
		return nil,"error in opening the file",err
	}
	defer file.Close()

	img,format,err := image.Decode(file)

	if err != nil{
		return nil,"error in decoding the file",err
	}

	return img,format,err

}

func convert(img image.Image, format string, outputFilePath string) error{

	outputFile, err := os.Create(outputFilePath)

	if err != nil{
		return err
	}

	var reqFormats = []string{"jpeg","png"}
	flag := false

	for _,req := range reqFormats{
		if req == format{
			flag = true
			break
		}
	}

	if !flag{
		return errors.New("invalid format")
	}

	switch format{
		case "jpeg":
			options := jpeg.Options{Quality: 100}
			err = jpeg.Encode(outputFile, img, &options)
		case "png":
			err = png.Encode(outputFile, img)
		default:
			err = errors.New("invalid format")
	}

	return err

}

func main(){
	
	outputFormat := "jpeg"
	filePath := ""
	
	if(len(os.Args ) == 0){
		fmt.Println("No file path provided")
	}
	
	for _,arg := range os.Args{
		if arg == "-h"{
			fmt.Println("Usage: go run main.go [file path] [-j/-p]")
		}else if arg == "-p"{
			outputFormat = "png"
		}else if arg == "-j"{
			outputFormat = "jpeg"
		}else{
			filePath = arg
		}
	}

	img,_,err := loadImage(filePath)
	
	if err != nil{
		fmt.Println(err)
		return
	}

	baseFileName := strings.TrimSuffix(filepath.Base(filePath),filepath.Ext(filePath))

	err = convert(img,outputFormat,baseFileName+"."+outputFormat)
	
	if err != nil{
		fmt.Println(err)
		return
	}else{
		fmt.Printf("Image successfully converted and saved as %s\n", baseFileName+"."+outputFormat)
	}
}