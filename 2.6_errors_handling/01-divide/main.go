package main

// Локально — полный файл для `go run .`.
// На Stepik отправляйте только func main() { ... }: пакет, импорты и divide там уже есть.

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	res, err := divide(a, b)
	if err != nil {
		fmt.Println("ошибка")
		return
	}
	fmt.Println(res)
}
