package main

import "fmt"

func process() (int, error) {
	// Выполнение некоторых операций
	return 42, nil
}

func main() {
	// Игнорируем оба возвращаемых значения
	process()

	// Используем только первое значение, ошибку игнорируем через _
	result, _ := process()
	fmt.Println(result)

	// Проверяем только ошибку
	_, err := process()
	if err == nil {
		fmt.Println("Ошибки нет")
	}
}
