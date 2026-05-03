package main

import "fmt"

func main() {
	rs := []rune("Это срез рун")

	for i := range rs {
		if rs[i] == 'р' {
			rs[i] = '*'
		}
	}

	fmt.Printf("Изменённый срез в виде строки: %s\n", string(rs))
}
