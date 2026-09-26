package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	inputFile := os.Args[1]
	fmt.Println("Reading from file:", inputFile)

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Process the file...
	lw := logWriter{}
	io.Copy(lw, file)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
