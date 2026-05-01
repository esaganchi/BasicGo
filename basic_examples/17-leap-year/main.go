package main

import "fmt"

func main() {
	// Високосный год: кратен 400 ИЛИ (кратен 4 и не кратен 100).
	var year int
	fmt.Scan(&year)
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
