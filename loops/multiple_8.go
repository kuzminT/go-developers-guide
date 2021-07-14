/**
Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
Программа в первой строке получает на вход число n - количество чисел в последовательности, во второй строке -- n чисел, входящих в данную последовательность.

Sample Input:

5
38 24 800 8 16
Sample Output:

40
**/

package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	var m int
	sum := 0

	for i := 1; i <= n; i++ {
		fmt.Scan(&m)
		if len(fmt.Sprint(m)) == 2 && m%8 == 0 {
			sum += m
		}
	}
	fmt.Println(sum)
}
