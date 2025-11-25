package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) // Stops the program with exit code 1
	}

	// // Original:
	// bs := make([]byte, 99999) // make, type of slice, number of elements in slice, create with n number of space
	// resp.Body.Read(bs)        // Create a byte slice, to be populated by the Read func

	// fmt.Println(string(bs))

	// Optimized: Not making byteSlice manually each time etc
	// Advanced
	io.Copy(os.Stdout, resp.Body) // Same output as before.
	// Using 'writer' to send data to the standard output, i.e. terminal, or write information to file persistance etc
	//

}
