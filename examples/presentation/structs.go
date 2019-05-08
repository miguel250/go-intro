package main

import "fmt"

type Animal struct {
	Name string
	Type string
}

func main() {
	a := Animal{
		Name: "Spot",
		Type: "Dog",
	}

	fmt.Println(a.Name) // Print animal name
}
