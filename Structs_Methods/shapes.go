package main

import "math"

type Shape interface {
	area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (triangle Triangle) area() float64 {
	return 0.5 * (triangle.Base * triangle.Height)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rectangle Rectangle) area() float64 {
	return rectangle.Width * rectangle.Height
}

type Circle struct {
	Radius float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
