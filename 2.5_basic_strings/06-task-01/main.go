// Задача 1: строка верна, если начинается с заглавной буквы и заканчивается точкой.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimRight(text, "\r\n")
	runes := []rune(text)
	if len(runes) == 0 {
		fmt.Println("Wrong")
		return
	}
	if unicode.IsUpper(runes[0]) && runes[len(runes)-1] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
