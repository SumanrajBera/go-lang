package main

import (
	"fmt"
	"github.com/SumanrajBera/reverseString"
	"github.com/common-nighthawk/go-figure"
	"github.com/google/uuid"
)

func main() {
	myFigure := figure.NewFigure("Hello World", "", true)
	myFigure.Print()

	id := uuid.New()
	fmt.Println(id)

	str := "Hello World"
	fmt.Println(reversestring.Reverse(str))
}
