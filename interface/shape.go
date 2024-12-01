package main

import "fmt"

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	area := t.height * t.base / 2
	return area
}
func (s square) getArea() float64 {
	area := s.sideLength * s.sideLength
	return area
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
