package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	inputPath := flag.String("input", "", "Path to the input file")

	flag.Parse()

	if *inputPath == "" {
		fmt.Println("Input path is required")
		os.Exit(1)
	}

	file, err := os.Open(*inputPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	io.Copy(os.Stdout, file)

}
