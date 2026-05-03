package main

import (
	"strings"
	"testing"
	"unicode"
)

// Та же проверка, что в main (тесты не могут вызывать main напрямую).
func right(s string) bool {
	s = strings.TrimRight(s, "\r\n")
	runes := []rune(s)
	if len(runes) == 0 {
		return false
	}
	return unicode.IsUpper(runes[0]) && runes[len(runes)-1] == '.'
}

func TestOk_stepikExamples(t *testing.T) {
	t.Run("как в примере — кириллица и точка", func(t *testing.T) {
		if !right("Быть или не быть.") {
			t.Fatal(`ожидали Right для "Быть или не быть."`)
		}
	})

	t.Run("первая буква строчная — Wrong", func(t *testing.T) {
		if right("быть или не быть.") {
			t.Fatal(`ожидали Wrong для строки со строчной первой буквой`)
		}
	})

	t.Run("нет точки в конце — Wrong", func(t *testing.T) {
		if right("Быть или не быть") {
			t.Fatal(`ожидали Wrong без точки в конце`)
		}
	})
}
