package main

import "fmt"

func main() {
	hello()
	multiply(3, 7)  // Результат: 21
	multiply(10, 5) // Результат: 50
}

func hello() {
	fmt.Println("Hello World")
}

func multiply(a int, b int) {
	result := a * b
	fmt.Println("Результат:", result)
}
