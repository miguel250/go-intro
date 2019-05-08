package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["dog"] = 6 // Add key and value to map

	fmt.Println(x)

	// Another way to create a map
	y := map[string]int{
		"key":  10,
		"key2": 20,
	}
	fmt.Println(y["key2"]) // access key
	delete(y, "key2")      // delete value from map

	// You can check if a key exist
	v, ok := y["key2"]
	fmt.Println(ok)
	fmt.Println(v)
}
