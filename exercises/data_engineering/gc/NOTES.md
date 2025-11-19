
- How to use comman line interface inputs in the script

* Require the "flag" package
* Using the flag.string to get a pointer to the input value
* Remember the pointer

```go
func main() {
	// 1. Define the flag
	// flag.String registers a new string flag with the command-line parser.
	// It takes three arguments:
	//   - The flag name ("input")
	//   - The default value ("") if the flag is not provided
	//   - A help message describing what the flag does
	// It returns a pointer to a string that will hold the flag's value.
	inputPathPtr := flag.String("input", "", "Path to the input file")

	// 2. Parse the command-line arguments
	// This examines the command-line arguments and assigns the values
	// to the variables registered with flag functions (like inputPathPtr).
	flag.Parse()

	// 3. Use the flag's value
	// We dereference the pointer (*) to get the actual string value.
	if *inputPathPtr == "" {
		fmt.Println("The --input flag is required.")
		fmt.Println("Usage: go run main.go --input \"/path/to/your/file\"")
	} else {
		fmt.Printf("The input path provided is: %s\n", *inputPathPtr)
	}
}
```
