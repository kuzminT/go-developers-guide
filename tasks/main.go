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
	printArea() float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) prinrArea() {
	fmt.Println(t.getArea())
}

func (s square) prinrArea() {
	fmt.Println(s.getArea())
}

func prinrArea(s shape) {
	s.printArea()
}

func main() {
	t := triangle{3, 4}
	s := square{3}
	t.prinrArea()
	s.prinrArea()
}
