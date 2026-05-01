package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}

// fibonacci: F(1)=1, F(2)=1, F(n)=F(n-1)+F(n-2) (Stepik).
func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	a, b := 1, 1
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
