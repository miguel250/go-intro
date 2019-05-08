package main

import "fmt"

type animal struct{}

func main() {
	var (
		i int
		f float64
		b bool
		s string
		p *animal
	)
	fmt.Printf("%v %v %v %q %v\n", i, f, b, s, p)
}
