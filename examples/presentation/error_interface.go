package main

import (
	"fmt"
)

// START OMIT
type MyError struct {
	Func    string
	Message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.Func, e.Message)
}

func Ops() error {
	return MyError{Func: "Ops", Message: "Failed!!!"}
}

func main() {
	err := Ops()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// END OMIT
