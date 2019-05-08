package main

import "fmt"

func main() {
	a := make([]string, 1, 10)
	fmt.Println(fmt.Sprintf("Len: %d cap: %d", len(a), cap(a)))

	b := make([]string, 0)
	fmt.Println(fmt.Sprintf("Len: %d cap: %d", len(b), cap(b)))

	b = append(b, "dog")
	fmt.Println(fmt.Sprintf("Len: %d cap: %d", len(b), cap(b)))

	b = append(b, "cat")
	fmt.Println(fmt.Sprintf("Len: %d cap: %d", len(b), cap(b)))
}
