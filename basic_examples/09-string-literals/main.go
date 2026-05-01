package main

import "fmt"

func main() {
	// 1) Строка в двойных кавычках:
	// перенос строки через \n, табуляция через \t
	s1 := "Привет\nМир\t!"

	// 2) Строка в обратных апострофах:
	// можно писать в несколько строк, \n и \t не интерпретируются
	s2 := `Привет
Мир\t!`

	fmt.Println("=== s1 (double quotes) ===")
	fmt.Println(s1)

	fmt.Println("=== s2 (backticks) ===")
	fmt.Println(s2)
}
