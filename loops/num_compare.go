/**
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.

Входные данные

Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.

Выходные данные

Программа должна вывести цифры, которые имеются в обоих числах, через пробел. Цифры выводятся в порядке их нахождения в первом числе.


Sample Input:

564 8954
Sample Output:

5 4
**/

package main

import "fmt"

func main() {
	var (
		a string
		b string
	)

	fmt.Scan(&a, &b)

	for _, char1 := range a {

		for _, char2 := range b {

			if char1 == char2 {
				fmt.Printf("%s ", string(char1))
			}
		}
	}
}