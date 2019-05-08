package main

import "fmt"

// x is copied and pased to the zero function
func zero(x int) {
	x = 0
}

func zeroPointer(x *int) {
	*x = 0
}

func main() {
	x := 1
	zero(x)
	fmt.Println(x)
	zeroPointer(&x) // & gets the address of the value
	fmt.Println(x)
}
