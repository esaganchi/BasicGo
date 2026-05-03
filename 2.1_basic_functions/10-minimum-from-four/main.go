package main

import "fmt"

func main() {
	fmt.Println(minimumFromFour())
}

// minimumFromFour читает четыре числа и возвращает минимальное (Stepik).
func minimumFromFour() int {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	if d < min {
		min = d
	}
	return min
}
