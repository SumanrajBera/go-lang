package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func checkType[T any](t T) {
	switch any(t).(type) {
	case int:
		fmt.Println("This is an interger")
	case string:
		fmt.Println("This is a string")
	case bool:
		fmt.Println("This is a boolean")
	case Shape:
		fmt.Println("This is a shape")
	default:
		fmt.Println("This is an unknown type")
	}
}

func main() {
	r := Rectangle{width: 24, height: 25}
	checkType(r)
	s := "Hello World"
	checkType(s)
	f := 0.5
	checkType(f)
}
