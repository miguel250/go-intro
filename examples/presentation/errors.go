package main

import (
	"errors"
	"fmt"
)

func Failed() error {
	return errors.New("Failed!!")
}

func main() {
	err := Failed()
	if err != nil {
		fmt.Println(err)
		return
	}
}
