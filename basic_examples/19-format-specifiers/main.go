package main

import (
	"fmt"
	"strings"
)

func main() {
	var x float64
	fmt.Scan(&x)

	if x > 10000 {
		fmt.Printf("%e\n", x)
		return
	}

	if x <= 0 {
		fmt.Printf("число %2.2f не подходит\n", x)
		return
	}

	squared := x * x
	s := fmt.Sprintf("%.10f", squared)
	dot := strings.Index(s, ".")
	fmt.Println(s[:dot+5])
}