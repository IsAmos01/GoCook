package main

import "fmt"

func main() {
	var x = 3.4
	// reflectType(x)
	reflectValue(&x)
	fmt.Println(x)
}
