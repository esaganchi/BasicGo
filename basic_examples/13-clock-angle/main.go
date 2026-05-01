package main

import "fmt"

func main() {
	// Часовая стрелка повернулась на d градусов с начала суток (12-часовой циферблат).
	// Найти целые часы h и целые минуты m.
	// За 12 часов стрелка проходит 360° → 30° в час, 0.5° в минуту.
	// d градусов = d/0.5 = 2d минут от полуночи.
	var d int
	fmt.Scan(&d)
	totalMinutes := 2 * d
	h := totalMinutes / 60
	m := totalMinutes % 60
	fmt.Printf("It is %d hours %d minutes.\n", h, m)
}
