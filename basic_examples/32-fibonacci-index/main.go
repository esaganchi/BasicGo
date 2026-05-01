package main

import "fmt"

func main() {
	// Дано A > 1. Найти n, что F_n = A (F_1=1, F_2=1, F_3=2, …), иначе -1.
	var a int
	fmt.Scan(&a)

	fPrev, fCurr := 1, 1
	n := 2

	for fCurr < a {
		fPrev, fCurr = fCurr, fPrev+fCurr
		n++
	}
	if fCurr == a {
		fmt.Println(n)
	} else {
		fmt.Println(-1)
	}
}
