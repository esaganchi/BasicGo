package main

import "fmt"

func main() {
	// Дано натуральное число N (≤ 10000). Вывести его последнюю цифру.
	var n int
	fmt.Scan(&n)
	fmt.Println(n % 10)
}
