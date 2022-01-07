package main

import "fmt"

func work(x int) int {
	return x * 10
}

func main() {
	m := make(map[int]int)
	var n int
	for i := 0; i < 10; i++ {
		fmt.Scan(&n)
		if _, ok := m[n]; !ok {
			m[n] = work(n)
		}
		fmt.Print(m[n])
	}
}
