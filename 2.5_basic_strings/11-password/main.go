package main

import "fmt"

func main() {
	var p string
	fmt.Scan(&p)
	if okPassword(p) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}

func okPassword(s string) bool {
	if len(s) < 5 {
		return false
	}
	for _, r := range s {
		digit := r >= '0' && r <= '9'
		lower := r >= 'a' && r <= 'z'
		upper := r >= 'A' && r <= 'Z'
		if !digit && !lower && !upper {
			return false
		}
	}
	return true
}
