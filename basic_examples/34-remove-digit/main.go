package main

import "fmt"

func main() {
	// Натуральное число (строкой) и цифра d — вывести число без всех вхождений этой цифры.
	var s string
	var d int
	fmt.Scan(&s, &d)

	target := '0' + rune(d)
	for _, ch := range s {
		if ch != target {
			fmt.Print(string(ch))
		}
	}
	fmt.Println()
}
