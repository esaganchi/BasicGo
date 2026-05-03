package main // исполняемая программа

import (
	"bufio"       // буферизованное чтение (строка целиком, не обрываясь на пробеле)
	"fmt"         // вывод Right / Wrong
	"os"          // стандартный ввод os.Stdin
	"strings"     // TrimRight, HasSuffix
	"unicode"     // IsUpper — заглавная ли буква (в т.ч. кириллица)
	"unicode/utf8" // корректно взять первый символ в UTF-8
)

func main() {
	// читаем всю строку до символа перевода строки (как в подсказке к задаче)
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// убираем только конец строки Windows/Mac/Unix, текст внутри не трогаем
	text = strings.TrimRight(text, "\r\n")

	// если строка удовлетворяет условиям — Right, иначе — Wrong
	if ok(text) {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}

// ok: true, если строка не пустая, начинается с заглавной буквы и заканчивается точкой
func ok(s string) bool {
	if s == "" { // пустая строка — не подходит
		return false
	}
	// первый символ строки в виде руны (Unicode-кодовой точки)
	first, _ := utf8.DecodeRuneInString(s)
	// битая кодировка или не заглавная буква — не подходит
	if first == utf8.RuneError || !unicode.IsUpper(first) {
		return false
	}
	// последний символ строки должен быть точкой '.'
	return strings.HasSuffix(s, ".")
}
