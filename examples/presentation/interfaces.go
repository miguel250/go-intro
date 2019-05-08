package main

import "fmt"

type Animal struct {
	Name string
	Type string
}

type Human struct {
	Name string
}

//START OMIT
type Interface interface {
	String() string
}

func (a Animal) String() string {
	return fmt.Sprintf("Hi, I am %s and I am a %s", a.Name, a.Type)
}

func (h Human) String() string {
	return fmt.Sprintf("Hi, I am %s and I am a human", h.Name)
}

func speak(i Interface) {
	fmt.Println(i.String())
}

func main() {
	h := Human{Name: "Miguel"}
	a := Animal{Name: "Spot", Type: "Dog"}

	speak(h)
	speak(a)
}

//END OMIT
