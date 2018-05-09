package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Tool for converting a file to a byte array.
// Usage: go run filetobytes.go <somefile.png>
func main() {
	path := os.Args[1]

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("unable to read file at %s: %v", path, err))
	}

	str := fmt.Sprintf("package files\nvar (\nsome_file = %#v\n)", data)

	err = ioutil.WriteFile("files.go", []byte(str), 0644)
	if err != nil {
		panic(fmt.Sprintf("unable to write file bytes: %v", err))
	}
}
