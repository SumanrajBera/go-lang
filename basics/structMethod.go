package main

import "fmt"

type rect struct {
	width  int
	height int
}

// (r rect) is receiver which means it belongs to type rect
func (r rect) area() int {
	return r.width * r.height
}

var r = rect{
	width: 10,
	height: 50,
}

func structMethod() {
	fmt.Println(r.area())
}
