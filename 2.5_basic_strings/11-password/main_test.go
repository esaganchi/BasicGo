package main

import "testing"

func TestOkPassword(t *testing.T) {
	if !okPassword("fdsghdfgjsdDD1") {
		t.Fatal("example")
	}
	if okPassword("abcd") {
		t.Fatal("short")
	}
	if okPassword("abc12!") {
		t.Fatal("bad char")
	}
	if !okPassword("abcde") {
		t.Fatal("min ok")
	}
}
