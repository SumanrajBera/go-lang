package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Describe() string {
	return fmt.Sprintf("My name is %v", a.Name)
}

type Dog struct {
	Animal
	Breed string
}

func (d Dog) Describe() string {
	return fmt.Sprintf("My name is %v and my breed is %v", d.Name, d.Breed)
}

func main() {
	a := Animal{Name: "Eagle"}
	fmt.Println(a.Describe())
	d := Dog{Breed: "Labrador", Animal: Animal{Name: "GoldyLocks"}}
	fmt.Println(d.Describe())
}
