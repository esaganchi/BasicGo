package main

import "fmt"

func main() {
	result1 := subtract(15, 8)
	result2 := subtract(50, 30)
	fmt.Println(result1)
	fmt.Println(result2)
}

func subtract(a, b int) int {
	return a - b
}
