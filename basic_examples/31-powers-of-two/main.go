package main

import "fmt"

func main() {
	// Все целые степени двойки ≤ N по возрастанию (1, 2, 4, …).
	var n int
	fmt.Scan(&n)

	p := 1
	first := true
	for p <= n {
		if !first {
			fmt.Print(" ")
		}
		fmt.Print(p)
		first = false
		p *= 2
	}
	fmt.Println()
}
