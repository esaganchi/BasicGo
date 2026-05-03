package main

import "fmt"

// func divide(a int, b int) int {
// 	return a / b
// }

// func main() {
// 	var input int
// 	fmt.Scan(&input)
// 	fmt.Println(divide(input, 5)) // Выведем результат
// }

func divide(a int, b int) int {
	return a / b
}

func main() {
	var input int
	_, err := fmt.Scan(&input) // fmt.Scan возвращает два параметра, первый — это количество считанных значений, второй — ошибка
	if err != nil {
		fmt.Println("Проверьте типы входных параметров.")
	} else {
		fmt.Println(divide(input, 5)) // Выводим результат, если ошибок нет
	}
}
