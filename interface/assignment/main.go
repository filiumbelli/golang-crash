package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: go run [go file name] [filename], Example: go run main.go my.txt")
		os.Exit(1)
	}
	filename := args[1]
	file, er1 := os.Open(filename)
	if er1 != nil {
		fmt.Printf("Error while opening %s: %v\n", filename, er1)
		os.Exit(1)
	}
	_, er2 := io.Copy(os.Stdout, file)
	if er2 != nil {
		fmt.Printf("Error while copying %s: %v\n", filename, er2)
	}

}
