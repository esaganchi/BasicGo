package main

import "fmt"

func main() {
	fmt.Println(divide(15, 5))               // Работает нормально
	fmt.Println(divide(4, 0))                // Генерирует ошибку
	fmt.Println("Program has been finished") // Не выполнится
}

func divide(x, y float64) float64 {
	if y == 0 {
		panic("division by zero!") // Генерация ошибки
	}
	return x / y
}
