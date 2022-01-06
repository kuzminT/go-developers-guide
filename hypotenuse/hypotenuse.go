package main

import (
	"fmt"
	"math"
)

// c=sqrt{{a^2}+{b^2}}

func main() {
	var (
		a,
		b float64
	)
	fmt.Scan(&a, &b)
	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	fmt.Println(c)
}
