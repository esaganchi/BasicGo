package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// printLenAndRunes печатает len в байтах и число рун для двух строк (как в исходном примере).
func printLenAndRunes(a, b string) {
	fmt.Println(len(a), len(b))
	fmt.Println(utf8.RuneCountInString(a), utf8.RuneCountInString(b))
}

// readPairsFromStdin читает из stdin пары строк: первая строка пары, вторая; пока есть ввод.
// Пустая первая строка пары — сигнал завершить ввод (вторая строка этой пары не читается).
func readPairsFromStdin() {
	in := bufio.NewScanner(os.Stdin)
	fmt.Println("Вводите по две строки подряд (каждая пара — с новой строки).")
	fmt.Println("Пустая строка в начале пары — выход. Конец файла (Ctrl+D) — тоже конец.")

	for {
		fmt.Print("строка 1> ")
		if !in.Scan() {
			if err := in.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "ошибка чтения первой строки:", err)
			}
			return
		}
		s1 := in.Text()
		if strings.TrimSpace(s1) == "" {
			return
		}

		fmt.Print("строка 2> ")
		if !in.Scan() {
			if err := in.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "ошибка чтения второй строки:", err)
			} else {
				fmt.Fprintln(os.Stderr, "вторая строка не введена: поток ввода закончился до ввода второй строки пары.")
			}
			return
		}
		s2 := in.Text()

		printLenAndRunes(s1, s2)
		fmt.Println()
	}
}

func main() {
	// Демо на фиксированных строках (как в учебнике).
	var en = "english"
	var ru = "русский"
	printLenAndRunes(en, ru)
	fmt.Println()

	readPairsFromStdin()
}
