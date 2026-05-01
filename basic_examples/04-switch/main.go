package main

import (
	"fmt"
	"time"
)

func main() {
	example5() // поменяйте на example1(), example2(), example3() или example4() для других примеров
}

// Пример 1: Switch по значению.
// Переменная сравнивается с константами в case (равенство).
// В одном case можно перечислить несколько значений через запятую.
// Подходит, когда нужно проверить конкретные значения (как несколько if x == 1, x == 2).
func example1() {
	switch day := time.Now().Weekday(); day {
	case time.Saturday, time.Sunday: // Несколько значений через запятую
		fmt.Println("Выходной")
	case time.Monday:
		fmt.Println("Тяжелый день")
	default: // Если ни один кейс не подошел
		fmt.Println("Обычный будний день")
	}
}

// Пример 2: Switch без выражения.
// После switch ничего нет, в каждом case — условие (как цепочка if / else if).
// Подходит для диапазонов и произвольных условий, а не только для равенства.
func example2() {
	var score int
	fmt.Print("Введите балл: ")
	fmt.Scanln(&score)
	switch {
	case score >= 90:
		fmt.Println("Отлично")
	case score >= 75:
		fmt.Println("Хорошо")
	default:
		fmt.Println("Нужно подтянуться")
	}
}

// Пример 4: Несколько значений в одном case (через запятую).
// В Go в одном case можно перечислить несколько значений — выполнится одна ветка для любого из них.
func example4() {
	var n int
	fmt.Print("Введите число 1-6: ")
	fmt.Scanln(&n)
	switch n {
	case 1, 2, 3:
		fmt.Println("Маленькое (1, 2 или 3)")
	case 4, 5, 6:
		fmt.Println("Среднее (4, 5 или 6)")
	default:
		fmt.Println("Другое число")
	}
}

// Пример 5: Калькулятор — switch по строке (оператор).
// Ввод: два числа и оператор (например: 10 5 +). В case "/" проверка деления на ноль.
func example5() {
	var a, b int
	var op string
	fmt.Print("Введите a, b и оператор (+, -, *, /): ")
	fmt.Scan(&a, &b, &op)

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("error")
		} else {
			fmt.Println(a / b)
		}
	default:
		fmt.Println("error")
	}
}

// Пример 3: Type switch.
// Специальный синтаксис для проверки типа значения (i.(type)).
// Используется с интерфейсами (any): в зависимости от реального типа выполняется разная ветка.
func example3() {
	identify(42)
	identify("привет")
	identify(3.14)
}

func identify(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Это число: %d\n", v)
	case string:
		fmt.Printf("Это строка длиной %d\n", len(v))
	default:
		fmt.Printf("Неизвестный тип: %T\n", v)
	}
}
