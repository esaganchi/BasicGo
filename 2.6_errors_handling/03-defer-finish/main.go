// package main

// import "fmt"

// func main() {
// 	defer finish() // Отложенный вызов функции finish
// 	fmt.Println("Program has been started")
// 	fmt.Println("Program is working")
// }

// func finish() {
// 	fmt.Println("Program has been finished")
// }

package main

import "fmt"

func main() {
	defer finish()                                // Этот вызов отложен
	defer fmt.Println("Program has been started") // Этот вызов отложен
	fmt.Println("Program is working")
}

func finish() {
	fmt.Println("Program has been finished")
}
