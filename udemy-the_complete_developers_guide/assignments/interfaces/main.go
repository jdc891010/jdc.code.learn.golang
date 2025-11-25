package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
	// printArea() float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(sh shape) { // shape is interface, and triangle and square is now a interface, and could be fed to the function.
	fmt.Println("Area:", sh.getArea())
}

func main() {
	t := triangle{height: 10, base: 5}
	printArea(t)

	s := square{sideLength: 5}
	printArea(s)
}
