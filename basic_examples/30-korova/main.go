package main

import "fmt"

func main() {
	// n коров на лугу: korova / korovy / korov (латиница, 0 < n < 100).
	var n int
	fmt.Scan(&n)

	var word string
	switch {
	case n%10 == 1 && n%100 != 11:
		word = "korova"
	case (n%10 == 2 || n%10 == 3 || n%10 == 4) && n%100 != 12 && n%100 != 13 && n%100 != 14:
		word = "korovy"
	default:
		word = "korov"
	}
	fmt.Printf("%d %s\n", n, word)
}
