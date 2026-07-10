package main
import "fmt"

func interpolation() {
	username := "sumanraj"
	age := 24
	weight := 84.226
	const favColor = "red"

	fmt.Printf("My name is %v \n", username)
	fmt.Printf("My age is %v \n", age)
	fmt.Printf("My weight is %1.f kg \n", weight)
	fmt.Printf("My favorite color is %v", favColor)
}
