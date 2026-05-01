package main

import "fmt"

func main() {
	num := 15
	fmt.Println("До вызова функции:", num)
	increase(num)
	fmt.Println("После вызова функции:", num)
}

func increase(n int) {
	fmt.Println("Внутри функции, до увеличения:", n)
	n += 10
	fmt.Println("Внутри функции, после увеличения:", n)
}
