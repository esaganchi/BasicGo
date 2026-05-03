package main

import "testing"

func TestFirstIndex(t *testing.T) {
	tests := []struct {
		x    string
		s    string
		want int
	}{
		{"awesome", "es", 2},
		{"hello", "ll", 2},
		{"hello", "xyz", -1},
		{"abc", "", 0},
		{"", "a", -1},
		{"aaaa", "aa", 0},
	}
	for _, tt := range tests {
		if got := firstIndex(tt.x, tt.s); got != tt.want {
			t.Errorf("firstIndex(%q, %q) = %d, хотели %d", tt.x, tt.s, got, tt.want)
		}
	}
}
