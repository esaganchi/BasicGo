package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Введите три числа через пробел: ")
	_, err := fmt.Scan(&a, &b, &c)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	result := multiplyAndAdd(a, b, c)
	fmt.Println(result)
}

func multiplyAndAdd(a, b, c int) int {
	return a*b + c
}
