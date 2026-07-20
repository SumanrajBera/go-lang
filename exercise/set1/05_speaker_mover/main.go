package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Mover interface {
	Move() string
}

type Functions interface {
	Speaker
	Mover
}

type Robot struct {
	name  string
	model string
}

type Dog struct {
	name  string
	breed string
}

func (r Robot) Speak() string {
	return fmt.Sprintf("%v-%v is speaking", r.name, r.model)
}

func (r Robot) Move() string {
	return fmt.Sprintf("%v-%v is moving", r.name, r.model)
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%v-%v is speaking", d.name, d.breed)
}

func (d Dog) Move() string {
	return fmt.Sprintf("%v-%v is moving", d.name, d.breed)
}

func Describe(f Functions) {
	fmt.Println(f.Speak())
	fmt.Println(f.Move())
}

func main() {
	r := Robot{name: "Chitti", model: "A65124"}
	d := Dog{name: "kuttuPullu", breed: "Golden Retriever"}
	Describe(r)
	Describe(d)
}
