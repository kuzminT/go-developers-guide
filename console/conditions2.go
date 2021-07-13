package main

import (
	"fmt"
	"runtime"
)

func main() {
	var a int
	fmt.Scan(&a)

	var (
		h int = a / 100
		d     = a % 100 / 10
		e     = a % 10
	)

	if h != d && d != e && e != h {
		fmt.Println("YES")

	} else {
		fmt.Println("NO")
	}

	// fmt.Print(h, d, e)
	fmt.Println(runtime.NumCPU())
}
