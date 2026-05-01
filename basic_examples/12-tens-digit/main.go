package main

import "fmt"

func main() {
	// Дано натуральное число (≤ 10000). Вывести число десятков (вторую цифру справа).
	var n int
	fmt.Scan(&n)
	fmt.Println((n / 10) % 10)
}
