package main

import "fmt"

/*

Go’s structs are typed collections of fields.
They’re useful for grouping data together to form records.
*/

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{"Alice", 31})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{"Ann", 40})

	s := person{"Sean", 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}