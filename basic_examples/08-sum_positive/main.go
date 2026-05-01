// Считываем целые числа до -1 или пока не введём 5 чисел. Суммируем только положительные.
// Запуск: go run .   Ввод: 3 5 -1   Вывод: 8   или: 1 2 3 4 5 → 15
package main

import "fmt"

func main() {
	const maxCount = 5 // прерываем цикл после 5 введённых значений
	sum := 0
	count := 0

	for {
		var n int
		fmt.Scan(&n)
		count++

		if n == -1 {
			break
		}
		if n > 0 {
			sum += n
		}
		if count >= maxCount {
			break
		}
	}
	fmt.Println(sum)
}
