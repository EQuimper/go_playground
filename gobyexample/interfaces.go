package main

import (
	"math"
	"fmt"
)

// Interfaces are named collections of method signatures.

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rectangle{3, 4}
	c := circle{5}

	measure(r)
	measure(c)
}
