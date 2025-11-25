package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// const ROOT_PATH string = "C:\\Users\\jdc\\Documents\\data\\gc_opendata\\batch" // Yeah, this is a local path, will need to change it.

func findFilesUsingGlob(rootDir string, pattern string) ([]string, error) {
	fsys := os.DirFS(rootDir)

	matches, err := fs.Glob(fsys, pattern)
	if err != nil {
		fmt.Printf("Error encountered: %v", err)

		return matches, err
	}

	if len(matches) == 0 {
		fmt.Println("No matches found, list is empty")
		os.Exit(1)
	} else {
		fmt.Println(matches)
	}

	return matches, err
}

func main() {
	inputPathPtr := flag.String("input", "", "Paath to the input file")

	flag.Parse()

	if *inputPathPtr == "" {
		fmt.Println("'--input' flag is required")
	} else {
		fmt.Printf("The input path provided is: %s", *inputPathPtr)
	}

	files, err := findFilesUsingGlob(*inputPathPtr, "*.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, file := range files {
		fullPath := filepath.Join(*inputPathPtr, file)
		fmt.Println(fullPath)

		readGCCsv(fullPath)
	}
}
