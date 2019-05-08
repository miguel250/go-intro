package main

import "fmt"

func main() {
	go func() {
		fmt.Println(1)
	}()

	fmt.Println(2)
}
