package main

import "fmt"

func main() {
	// Наибольшее число на отрезке [a, b], кратное 7; если такого нет — NO (Stepik 1.13).
	var a, b int
	fmt.Scan(&a, &b)

	// Наибольшее число ≤ b, кратное 7. Для b < 0 нельзя писать (b/7)*7: в Go деление
	// целых идёт к нулю, и получится значение больше b (например b=-8 → -7).
	q := b / 7
	if b < 0 && b%7 != 0 {
		q--
	}
	k := q * 7
	if k >= a {
		fmt.Println(k)
	} else {
		fmt.Println("NO")
	}
}
