package main

import "fmt"

type Animal struct {
	Name string
	Type string
}

func (a Animal) String() string {
	return fmt.Sprintf("Hi, I am %s name %s", a.Name, a.Type)
}

func main() {
	a := Animal{
		Name: "Spot",
		Type: "Dog",
	}

	fmt.Println(a.String())
	fmt.Println(a) // fmt will use call "String" if it exists in a struct
}
