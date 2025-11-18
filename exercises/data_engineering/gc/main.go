package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const ROOT_PATH string = "C:\\Users\\jdc\\Documents\\data\\gc_opendata\\batch"

func findFilesUsingGlob(rootDir string, pattern string) ([]string, error) {
	fsys := os.DirFS(rootDir)

	matches, err := fs.Glob(fsys, pattern)
	if err != nil {
		fmt.Printf("Error encountered: %v", err)

		return matches, err
	}

	if len(matches) == 0 {
		fmt.Println("No matches found, list is empty")
	} else {
		fmt.Println(matches)
	}

	return matches, err
}

func main() {
	files, err := findFilesUsingGlob(ROOT_PATH, "*/*.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, file := range files[:10] {
		fullPath := filepath.Join(ROOT_PATH, file)
		fmt.Println(fullPath)

		readGCCsv(fullPath)
	}
}
