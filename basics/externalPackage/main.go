package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/google/uuid"
	"fmt"
)

func main() {
	myFigure := figure.NewFigure("Hello World", "", true)
	myFigure.Print()

	id := uuid.New()
	fmt.Println(id)
}
