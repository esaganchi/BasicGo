package main

import "fmt"

func main() {
	multiply(4, 7)  // Результат: 21
	multiply(10, 5) // Результат: 50
}

func multiply(a int, b int) {
	result := a * b
	fmt.Println("Результат:", result)
}
