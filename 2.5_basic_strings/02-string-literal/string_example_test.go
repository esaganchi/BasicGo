package main

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

// Example_basicString — демонстрация вывода; эталон в комментарии // Output:.
func Example_basicString() {
	var s string = "Это строка"

	fmt.Printf("Длина строки: %d байт\n", len(s))
	fmt.Printf("Напечатаем только второе слово в кавычках: \"%v\"\n", s[7:])

	s = s + " Новая строка"
	fmt.Printf("%v\n", s)

	for _, b := range s {
		fmt.Printf("%v ", b)
	}
	fmt.Print("\n")

	// Output:
	// Длина строки: 19 байт
	// Напечатаем только второе слово в кавычках: "строка"
	// Это строка Новая строка
	// 1069 1090 1086 32 1089 1090 1088 1086 1082 1072 32 1053 1086 1074 1072 1103 32 1089 1090 1088 1086 1082 1072
}

// TestCyrillicStringLiteral — обычные проверки через testing.T (assert-стиль вручную).
func TestCyrillicStringLiteral(t *testing.T) {
	s := "Это строка"

	t.Run("len считает байты UTF-8", func(t *testing.T) {
		if got, want := len(s), 19; got != want {
			t.Errorf("len(s) = %d, хотели %d", got, want)
		}
	})

	t.Run("utf8.RuneCountInString считает символы", func(t *testing.T) {
		if got, want := utf8.RuneCountInString(s), 10; got != want {
			t.Errorf("RuneCountInString(s) = %d, хотели %d (10 символов фразы)", got, want)
		}
	})

	t.Run("срез по байтам s[7:] — второе слово", func(t *testing.T) {
		got := s[7:]
		want := "строка"
		if got != want {
			t.Errorf("s[7:] = %q, хотели %q", got, want)
		}
	})

	t.Run("конкатенация даёт новую строку", func(t *testing.T) {
		got := s + " Новая строка"
		want := "Это строка Новая строка"
		if got != want {
			t.Errorf("получили %q, хотели %q", got, want)
		}
	})

	t.Run("range идёт по рунам", func(t *testing.T) {
		extended := s + " Новая строка"
		var runes []rune
		for _, r := range extended {
			runes = append(runes, r)
		}
		if got, want := len(runes), 23; got != want {
			t.Errorf("число рун после конкатенации = %d, хотели %d (10 + 13 символов суффикса)", got, want)
		}
		if got, want := string(runes), extended; got != want {
			t.Errorf("сборка рун обратно в строку не совпала")
		}
	})
}
