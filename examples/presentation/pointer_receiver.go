package main

import "fmt"

type Animal struct {
	Name string
	Type string
}

func (a Animal) String() string {
	return fmt.Sprintf("Hi, I am %s and I am a %s", a.Name, a.Type)
}

// START OMIT

// Pointer receiver
func (a *Animal) SetName(n string) {
	a.Name = n
}

// Value receiver
func (a Animal) setType(t string) {
	a.Type = t
}

func main() {
	a := Animal{Name: "Spot", Type: "Dog"}
	a.SetName("Lighting")
	a.setType("cat")
	fmt.Println(a)
}

// END OMIT
