package main

import "fmt"

func main() {
	// Дано натуральное число (≤ 10000). Вывести первую (левую) цифру.
	var n int
	fmt.Scan(&n)
	for n >= 10 {
		n = n / 10
	}
	fmt.Println(n)
}
