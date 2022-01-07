package main

import (
	"fmt"
	"strconv"
)

func adding(x string, y string) int64 {
	removeTrash := func(s string) string {
		var newStringB []rune
		for _, v := range s {
			if _, err := strconv.Atoi(string(v)); err != nil {
				continue
			}
			newStringB = append(newStringB, v)
		}
		return string(newStringB)
	}

	var sum int64
	for _, n := range [2]string{x, y} {
		x = removeTrash(n)
		b, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			panic(err)
		}
		sum += b
	}

	return sum
}

func main() { //nolint
	x, y := "%^80", "hhhhh20&&&&nd"
	fmt.Println(adding(x, y))

}
