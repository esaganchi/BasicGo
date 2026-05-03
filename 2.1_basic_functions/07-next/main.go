package main // исполняемая программа — пакет main

import "fmt" // ввод-вывод: Print, Scan и т.д.

func main() {
	var a, b int                   // два целых числа со stdin
	var firstName, lastName string // имя и фамилия (два слова)

	fmt.Print("Введите два числа и имя с фамилией через пробел: ")
	// Scan читает значения и записывает по адресам &...; _ — число успешно считанных полей (не используем)
	_, err := fmt.Scan(&a, &b, &firstName, &lastName)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return // без корректного ввода дальше не идём
	}

	// calculate возвращает два значения: сумму чисел и полное имя
	sum, fullName := calculate(a, b, firstName, lastName)
	fmt.Println(sum)
	fmt.Println(fullName)
}

// calculate: a+b и строка "имя фамилия"; возвращает (int, string)
func calculate(a, b int, firstName, lastName string) (int, string) {
	total := a + b
	fullName := firstName + " " + lastName
	return total, fullName
}
