package main

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func Example_lenVsRune() {
	var en = "english"
	var ru = "русский"
	fmt.Println(len(en), len(ru))
	fmt.Println(utf8.RuneCountInString(en), utf8.RuneCountInString(ru))

	// Output:
	// 7 14
	// 7 7
}

func TestLenAndRunePairs(t *testing.T) {
	en, ru := "english", "русский"

	t.Run("len считает байты UTF-8", func(t *testing.T) {
		if got, want := len(en), 7; got != want {
			t.Errorf("len(%q) = %d, хотели %d", en, got, want)
		}
		if got, want := len(ru), 14; got != want {
			t.Errorf("len(%q) = %d, хотели %d (7 букв × 2 байта в UTF-8)", ru, got, want)
		}
	})

	t.Run("RuneCountInString считает символы", func(t *testing.T) {
		if got, want := utf8.RuneCountInString(en), 7; got != want {
			t.Errorf("RuneCountInString(%q) = %d, хотели %d", en, got, want)
		}
		if got, want := utf8.RuneCountInString(ru), 7; got != want {
			t.Errorf("RuneCountInString(%q) = %d, хотели %d", ru, got, want)
		}
	})

	t.Run("для кириллицы len больше числа рун", func(t *testing.T) {
		if got := len(ru); got <= utf8.RuneCountInString(ru) {
			t.Errorf("ожидали len(%q) > числа рун, получили len=%d, рун=%d", ru, got, utf8.RuneCountInString(ru))
		}
	})
}
