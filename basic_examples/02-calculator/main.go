package main

import "fmt"

func main() {
	main2()
}

func main1() {
	// Константа НДС (20%)
	const VAT = 0.20
	var price float64

	// Читаем цену товара
	fmt.Scanln(&price)

	// Вычисляем цену с НДС: цена + (цена * НДС) = цена * (1 + НДС)
	total := price * (1 + VAT)

	// Выводим результат с 2 знаками после запятой
	fmt.Printf("total=%.2f", total)
}

func main2() {
	var a, b int
	fmt.Print("Введите два числа через пробел: ")
	fmt.Scanf("%d %d", &a, &b)

	// Считайте два целых числа a и b.
	// Поменяйте их местами, используя промежуточную переменную (можно через :=).

	// напишите свой код тут
	tmp := a
	a = b
	b = tmp

	fmt.Printf("%d %d\n", a, b)
}
