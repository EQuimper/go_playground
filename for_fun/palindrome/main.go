package main

import (
	"fmt"
	"strings"
)

func reverse(a []string) []string {
	var arr []string

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]

		arr = a
	}

	return arr
}

func checkIfPalindrome(s string) bool {
	var reverseArr []string

	arr := strings.Split(s, ",")

	reverseArr = reverse(arr)

	str := strings.Join(reverseArr[:], ",")

	if str != s {
		return false
	}

	return true
}

func main() {
	c := checkIfPalindrome("A but tuba")

	fmt.Println("This is a palindrome", c)
}
