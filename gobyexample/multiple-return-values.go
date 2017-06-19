package main

import "fmt"

// Function gonna return 2 int
func vals() (int, int) {
	return 3, 7
}

func main() {
	// We destruc here both values
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	/*
		If you only want a subset of the returned values,
		use the blank identifier _.
	*/
	_, c := vals()
	fmt.Println(c)
}
