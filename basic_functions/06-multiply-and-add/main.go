package main

import "fmt"

func main() {
	result := multiplyAndAdd(5, 3, 2)
	fmt.Println(result)
}

func multiplyAndAdd(a, b, c int) int {
	return a*b + c
}
