package main

import "testing"

func TestOddIndexChars(t *testing.T) {
	if got, want := oddIndexChars("ihgewlqlkot"), "hello"; got != want {
		t.Errorf("oddIndexChars(stepik) = %q, хотели %q", got, want)
	}
	if got, want := oddIndexChars("abcdef"), "bdf"; got != want {
		t.Errorf("oddIndexChars(%q) = %q, хотели %q", "abcdef", got, want)
	}
	if got, want := oddIndexChars(""), ""; got != want {
		t.Errorf("пустая строка: got %q", got)
	}
	if got, want := oddIndexChars("a"), ""; got != want {
		t.Errorf("одна буква: got %q", got)
	}
}
