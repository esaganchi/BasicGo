// Первое вхождение подстроки: две строки X и S (каждая с новой строки); индекс в X или -1.
package main // исполняемая программа

import (
	"bufio"   // построчное чтение из stdin
	"fmt"     // вывод индекса
	"os"      // stdin
	"strings" // strings.Index — стандартный поиск подстроки
)

func main() { // точка входа
	in := bufio.NewScanner(os.Stdin) // сканер читает ввод по строкам

	in.Scan()      // первая строка — строка X
	x := in.Text() // без символа конца строки

	in.Scan()      // вторая строка — подстрока S
	s := in.Text() // ищем её внутри X

	fmt.Println(firstIndex(x, s)) // индекс первого вхождения S в X; если нет — -1
}

// firstIndex — индекс первого вхождения s в x, либо -1 (обёртка над strings.Index для тестов и main).
func firstIndex(x, s string) int {
	return strings.Index(x, s)
}
