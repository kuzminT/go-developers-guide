package main

//Дано трехзначное число. Найдите сумму его цифр.
//
//Формат входных данных
//На вход дается трехзначное число.
//
//Формат выходных данных
//Выведите одно целое число - сумму цифр введенного числа.
//
//Sample Input:
//
//745
//Sample Output:
//
//16

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		 num string
		 sum int
	)

	fmt.Scan(&num)
	var a int
	for _, c := range num {
		a, _ = strconv.Atoi(string(c))
		sum += a
	}
	fmt.Println(sum)
}
