package main

import "fmt"

func main() {
	var a, b int64
	fmt.Scan(&a, &b)
	sum := a + b
	if sum%2 == 0 {
		fmt.Println(sum / 2)
	} else {
		fmt.Printf("%.1f\n", float64(sum)/2)
	}
}
