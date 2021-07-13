package main

import "fmt"

func getSumOfNumber(num int) int {
	s := 0
	s += num / 100
	s += num % 100 / 10
	s += num % 10
	return s
}

func main() {
	var a int
	fmt.Scan(&a)

	var (
		firstNum  int = a / 1000
		secondNum int = a % 1000
	)

	s1 := getSumOfNumber(secondNum)
	s2 := getSumOfNumber(firstNum)

	if s1 == s2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
