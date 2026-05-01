package main

import "fmt"

func main() {
	main3()
}

// main1: пример оператора if (один вариант)
func main1() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scanln(&n)

	if n > 0 {
		fmt.Println("Число положительное")
	}
}

// main2: пример оператора if-else (два варианта)
func main2() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scanln(&n)

	if n > 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Число не положительное (ноль или отрицательное)")
	}
}

// main3: пример if-else if-else (несколько условий)
func main3() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scanln(&n)

	if n > 0 {
		fmt.Println("Положительное")
	} else if n < 0 {
		fmt.Println("Отрицательное")
	} else {
		fmt.Println("Ноль")
	}
}
