//Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.
//
//Пример вызова вашей функции:
//
//a, b := sumInt(1, 0)
//fmt.Println(a, b)
//
//Результат: 2, 1
//
//Sample Input:
//
//Sample Output:

package main

import "fmt"

func main() {
	//fmt.Scan(&)
	fmt.Println(sumInt(2, 1, 4))
}

func sumInt(args ...int) (int, int) {
	var sum int
	for _, n := range args {
		sum += n
	}
	return len(args), sum
}
