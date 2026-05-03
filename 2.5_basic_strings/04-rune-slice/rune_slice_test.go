package main

import "fmt"

// Example_rune — как ExampleRune в учебнике; суффикс _rune, чтобы не требовать тип Rune в пакете.
func Example_rune() {
	rs := []rune("Это срез рун")

	for i := range rs {
		if rs[i] == 'р' {
			rs[i] = '*'
		}
	}

	fmt.Printf("Изменённый срез в виде строки: %s\n", string(rs))

	// Output:
	// Изменённый срез в виде строки: Это с*ез *ун
}
