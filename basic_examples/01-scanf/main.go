package main

import "fmt"

func main() {
	var day int
	var month string
	fmt.Scanf("%d %s", &day, &month)
	// если ввести "12 March" → day=12, month="March"
	fmt.Printf("День: %d, Месяц: %s\n", day, month)
}
