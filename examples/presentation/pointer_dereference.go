package main

import "fmt"

type Animal struct {
	Name string
	Type string
}

func NameAnimal(a *Animal) {
	a.Name = "Lighting"
}

func main() {
	a := Animal{
		Name: "Spot",
		Type: "Dog",
	}

	NameAnimal(&a)
	fmt.Println(a.Name)
}
