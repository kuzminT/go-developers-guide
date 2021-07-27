package main

//Напишите программу, принимающая на вход число N (N \geq 4)N(N≥4), а затем N целых чисел-элементов среза.
//На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
//Sample Input:
//5
//41 -231 24 49 6
//Sample Output:
//
//49

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sl :=  make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&sl[i])
	}
	fmt.Println(sl[3])
}