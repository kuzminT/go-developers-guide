package main

import (
	"fmt"
	"strconv"
)

func main() {

	fn := func(input uint) uint {

		strInput := fmt.Sprint(input)

		var strResult string
		for _, n := range strInput {
			nInt, _ := strconv.ParseUint(string(n), 10, 32)
			if nInt != 0 && nInt%2 == 0 {
				strResult += string(n)
			}
		}
		result, _ := strconv.ParseUint(strResult, 10, 32)

		if result == 0 {
			return 100
		}

		return uint(result)
	}

	fmt.Println(fn(345))
}
