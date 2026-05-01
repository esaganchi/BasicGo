package main

import "fmt"

// Пример для go test: игнорирование возвращаемых значений и _.
func Example() {
	// Игнорируем оба возвращаемых значения
	process()

	// Используем только одно из значений
	result, _ := process()
	fmt.Println(result)

	// Проверяем только ошибку
	_, err := process()
	if err == nil {
		fmt.Println("Ошибки нет")
	}

	// Output:
	// 42
	// Ошибки нет
}
