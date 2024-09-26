package main

import (
	"flag"
	"fmt"
	"jsonBin/files"
	"jsonBin/storage"
)

func main() {
	filePath := flag.String("file", "", "Path to JSON file")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("You need to pass the path to the file using the --file flag")
	}

	if !files.IsJSONFile(*filePath) {
		fmt.Println("File don't have extension .json")
	}

	data, err := files.ReadFile(*filePath)
	if err != nil {
		fmt.Printf("Error for reading file: %v", err)
	}

	binList, err := storage.LoadBinListBytes(data)
	if err != nil {
		fmt.Printf("Error for unmarshal JSON: %v", err)
	}
	fmt.Printf("%+v", binList)
}
