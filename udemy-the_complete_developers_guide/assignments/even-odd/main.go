package main

import "fmt"

// section 3: Assignment: Even or Odd?
func main() {
	src_numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Creating a slice, since we didn't set a predefined size, although in this case, we could've.
	// could be src_numbers := [10]unit{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, since we know the values will be < 255, hence uint8,

	for _, v := range src_numbers {
		// remember, no need for "in" keyword, just := to set the value in the loop/range.
		//Index is first return, not utilized, default to "_".
		if v%2 == 0 { // remainer to determine odd and even, is easiest
			fmt.Printf("%d is even \n", v) // format print to include variable number
		} else {
			fmt.Printf("%d is odd \n", v)
		}
	}
}
