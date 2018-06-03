package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	outputFile = "file.go"
)

// Tool for converting a file to a byte array.
// Usage: go run main.go <somefile.png>
func main() {
	path := os.Args[1]

	log.Printf("File path: %s", path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	str := fmt.Sprintf("package files\nvar (\nsome_file = %#v\n)", data)

	err = ioutil.WriteFile(outputFile, []byte(str), 0644)
	if err != nil {
		log.Fatalf("Unable to write file bytes: %v", err)
	}

	log.Printf("Converted file to bytes and saved as %s", outputFile)
}
