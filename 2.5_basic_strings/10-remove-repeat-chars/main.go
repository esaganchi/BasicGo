package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	fmt.Println(keepCharsAppearingOnce(a))
}

func keepCharsAppearingOnce(s string) string {
	var b strings.Builder
	for _, ch := range s {
		if strings.Count(s, string(ch)) == 1 {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
