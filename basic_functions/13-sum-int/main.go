package main

import "fmt"

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}

// sumInt возвращает количество аргументов и их сумму (Stepik, variadic ...int).
func sumInt(nums ...int) (int, int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return len(nums), sum
}
