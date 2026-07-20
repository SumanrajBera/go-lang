package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.radius
}

func main() {
	rect := Rectangle{
		20.5, 30.2,
	}

	circ := Circle {
		20.54,
	}

	fmt.Printf("Area of rectangle: %.2f \n", rect.Area())
	fmt.Printf("Perimeter of rectangle: %.2f \n", rect.Perimeter())
	fmt.Printf("Area of Circle: %.2f \n", circ.Area())
	fmt.Printf("Perimeter of Circle: %.2f \n", circ.Perimeter())
}
