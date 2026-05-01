package ask

import "fmt"

// AskAndGreet запрашивает имя и возраст с ввода и выводит приветствие.
func AskAndGreet() {
	var name string
	var age int

	fmt.Print("Введите имя: ")
	fmt.Scan(&name)

	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	fmt.Printf("Привет, %s! Тебе %d лет.\n", name, age)
}
