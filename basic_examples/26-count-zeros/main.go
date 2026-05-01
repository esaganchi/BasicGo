package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	count := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		if x == 0 {
			count++
		}
	}
	fmt.Println(count)
}
