package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println(vote(x, y, z))
}

// vote возвращает 0 или 1, который встречается среди x, y, z чаще (Stepik).
func vote(x, y, z int) int {
	s := x + y + z
	if s >= 2 {
		return 1
	}
	return 0
}
