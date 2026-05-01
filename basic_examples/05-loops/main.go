package main

import "fmt"

func main() {
	example1()
	example2()
	example3()
	example4()
	example5()
	example6()
	example7()
}

// Классический for: инициализация; условие; шаг
func example1() {
	for i := 0; i < 5; i++ {
		fmt.Println("Шаг", i)
	}
}

// В Go нет ключевого слова while. Если в for оставить только условие — работает как while в других языках.
func example2() {
	// Как "while (n < 100)" в C/Java/JS:
	n := 1
	for n < 100 {
		n *= 2 // Удваиваем n, пока оно не станет >= 100
	}
	fmt.Println("n =", n)

	// Ещё пример: "while (x < 5)" — печатаем 0, 1, 2, 3, 4
	x := 0
	for x < 5 {
		fmt.Println("x =", x)
		x++
	}
}

// Бесконечный for: без условия. Чтобы не зависнуть навсегда, внутри обычно ставят условие и break
func example3() {
	count := 0
	for {
		fmt.Println("Я работаю...")
		// Чтобы не зависнуть навсегда, внутри обычно ставят условие и break
		count++
		if count >= 3 {
			break
		}
	}
}

// continue пропускает итерацию, break выходит из цикла. Вывод: 1, 2, 3, 4, 6, 7
func example4() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue // Число 5 будет пропущено
		}
		if i == 8 {
			break // На числе 8 цикл полностью остановится
		}
		fmt.Println(i)
	}
}

// Проверка числа на простоту: цикл по возможным делителям, break при первом найденном.
func example5() {
	var n int
	fmt.Scan(&n)

	// Число 1 не является простым по определению; отрицательные и 0 — тоже не простые
	if n < 2 {
		fmt.Println("no")
		return
	}

	// Предполагаем, что число простое, пока не найдём делитель
	isPrime := true

	// Проверяем делители от 2 до √n: если есть делитель, то один из них не больше корня из n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			isPrime = false
			break // Нашли делитель — число составное, дальше проверять не нужно
		}
	}

	if isPrime {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

// Поиск минимума и максимума среди n введённых чисел.
func example6() {
	var n int
	if _, err := fmt.Scan(&n); err != nil || n <= 0 {
		return // Если n <= 0, ничего не делаем
	}

	var current int
	fmt.Scan(&current) // Считываем первое число, чтобы инициализировать min и max

	min := current
	max := current

	// Начинаем со второго числа, так как первое уже считали
	for i := 1; i < n; i++ {
		fmt.Scan(&current)
		if current < min {
			min = current
		}
		if current > max {
			max = current
		}
	}

	fmt.Printf("min=%d max=%d\n", min, max)
}

// Считываем целые числа до -1, суммируем только положительные. Бесконечный цикл и break.
func example7() {
	sum := 0
	for {
		var n int
		fmt.Scan(&n)
		if n == -1 {
			break
		}
		if n > 0 {
			sum += n
		}
	}
	fmt.Println(sum)
}
