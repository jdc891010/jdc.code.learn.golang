package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const S3LINK string = "http://goldencheetah-opendata.s3-website-us-east-1.amazonaws.com/data/"

var userId string = "000db8a2-a1f6-42bd-8228-fdfae659f476"

func main() {
	out, err := os.Create("tmp.zip") // Creating a tmp file and opening it to allow interaction
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	resp, err := http.Get(S3LINK + userId + ".zip") // Get / retrieve the actual data
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close() // pushed function into a list, list of saved calls is executed after the surrounding function returns

	// Copying the content of the get request,to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Close the file so we can open it again
	out.Close()

	fmt.Println("File downloaded successfully!") //

	// using = instead of :=, since no NEW variable, err was previously defined
	err = os.Mkdir(userId, 0755) // 0755 is the permissions code -> The owner has full permissions: read, write, and execute.
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unzip
	r, err := zip.OpenReader("tmp.zip")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, f := range r.File { // The `for` loop iterates through each file in the zip archive.
		fmt.Printf("Extracting %s\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		out, err := os.Create(filepath.Join(userId, f.Name)) // Creating the file on disk, then copying content into the file.
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		_, err = io.Copy(out, rc)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		rc.Close()
		out.Close()
	}
	r.Close()

	err = os.Remove("tmp.zip") // Ensure the file is closed, before removing it
	if err != nil {
		fmt.Println("Error removing tmp.zip:", err)
		os.Exit(1)
	}
	fmt.Println("Removed tmp.zip")

}
