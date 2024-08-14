package main

import (
	"fmt"
	"os"
)

func main(){
	if(len(os.Args ) == 0){
		fmt.Println("No file path provided")
	}
	filePath := "./" + os.Args[1]

	content,err := os.ReadFile(filePath)

	if(err != nil){
		fmt.Println("Error reading file")
		return
	}

	fmt.Println(string(content))
}