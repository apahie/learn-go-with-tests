package main

import "math"

type Rectangle struct {
	width  float64
	height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
