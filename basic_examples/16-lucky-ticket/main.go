package main

import "fmt"

func main() {
	// Счастливый билет: сумма первых трёх цифр = сумма последних трёх.
	var n int
	fmt.Scan(&n)
	first := n/100000 + (n/10000)%10 + (n/1000)%10   // сумма 1–3 цифр
	last := (n/100)%10 + (n/10)%10 + n%10             // сумма 4–6 цифр
	if first == last {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
