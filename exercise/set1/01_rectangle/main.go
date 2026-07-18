package main

import "fmt"

type rectangleMeasurement interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) scale(factor float64) rectangle {
	newWidth := r.width * factor
	newHeight := r.height * factor

	return rectangle{
		width:  newWidth,
		height: newHeight,
	}
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	r := rectangle {
		width: 10,
		height: 20,
	}

	fmt.Println("Area of rectangle:", r.area())
	fmt.Println("Perimeter of rectangle:", r.perimeter())

	newReactangle := r.scale(50)
	fmt.Printf("Width of new rectangle: %v \nHeight of new rectangle: %v ",newReactangle.width, newReactangle.height)
}
