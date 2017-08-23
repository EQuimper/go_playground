package main

import (
	"fmt"
)

type numbers []int

func newNumbersList(num int) numbers {
	list := numbers{}

	for i := 0; i < num; i++ {
		list = append(list, i)
	}

	return list
}

func checkIfOddOrEven(num int) {
	if num%2 == 0 {
		fmt.Printf("%v is even \n", num)
	} else {

		fmt.Printf("%v is odd \n", num)
	}
}

func (n numbers) loopOver() {
	for _, num := range n {
		checkIfOddOrEven(num)
	}
}

func main() {
	nums := newNumbersList(11)

	nums.loopOver()
}
