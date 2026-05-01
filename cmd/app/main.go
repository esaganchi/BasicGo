package main

import (
	"fmt"
	"time"

	"basicgo/internal/ask"
	"basicgo/internal/hello"
)

func main() {
	fmt.Println(hello.Greet())
	fmt.Print(hello.PersonInfo("Alice", 25))
	fmt.Println("Сейчас:", time.Now().Format(time.RFC3339))

	ask.AskAndGreet()
}
