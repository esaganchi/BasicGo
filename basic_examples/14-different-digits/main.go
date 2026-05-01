package main

import "fmt"

func main() {
	// Дано трёхзначное число. Все ли цифры различны?
	var n int
	fmt.Scan(&n)
	a := n / 100       // первая цифра (сотни)
	b := (n / 10) % 10 // вторая цифра (десятки)
	c := n % 10        // третья цифра (единицы)
	if a != b && b != c && a != c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
