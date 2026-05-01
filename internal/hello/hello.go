package hello

import "fmt"

// Greet возвращает приветствие.
func Greet() string {
	return "Hello, Go!"
}

// PersonInfo возвращает строку с именем и возрастом.
func PersonInfo(name string, age int) string {
	return fmt.Sprintf("Имя: %s, Возраст: %d\n", name, age)
}
