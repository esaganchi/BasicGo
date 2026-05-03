package main

import "testing"

func TestKeepCharsAppearingOnce(t *testing.T) {
	if got, want := keepCharsAppearingOnce("zaabcbd"), "zcd"; got != want {
		t.Errorf("stepik: got %q, want %q", got, want)
	}
	if got, want := keepCharsAppearingOnce("aabb"), ""; got != want {
		t.Errorf("all repeated: got %q", got)
	}
	if got, want := keepCharsAppearingOnce("abc"), "abc"; got != want {
		t.Errorf("all unique: got %q", got)
	}
}
