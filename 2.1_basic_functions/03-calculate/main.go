package main

import "fmt"

func main() {
	calculate(2, 4, 1.5, 2.5, 3.0)
}

func calculate(x, y int, p, q, r float64) {
	sumInt := x + y
	sumFloat := p + q + r
	fmt.Println("Сумма целых чисел:", sumInt)
	fmt.Println("Сумма дробных чисел:", sumFloat)
}
