package main

import "fmt"

// func main() {
// 	var name string
// 	var age int
// 	fmt.Print("Введите имя: ")
// 	fmt.Scan(&name)
// 	fmt.Print("Введите возраст: ")
// 	fmt.Scan(&age)

// 	fmt.Println(name, age)
// }

// func main() {
// 	var a int
// 	fmt.Scan(&a)
// 	var b int = a*2 + 100
// 	fmt.Println(b)
// }

// func main() {
// 	var a int
// 	fmt.Scan(&a)
// 	aString := fmt.Sprint(a)
// 	fmt.Println(string(aString[len(aString)-2]))
// }

func main() {
	var a int
	fmt.Scan(&a)
	minutes := 2 * (a % 30)
	hours := a / 30
	fmt.Printf("It is %d hours %d minutes.\n", hours, minutes)
}
