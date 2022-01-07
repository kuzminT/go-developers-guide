package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Trim(str, "\n")
	vals := strings.Split(str, ";")
	nums := make([]float64, len(vals))
	for i, v := range vals {
		v = strings.ReplaceAll(strings.ReplaceAll(v, ",", "."), " ", "")
		newV, err := strconv.ParseFloat(v, 64)
		if err != nil {
			panic(err)
		}
		nums[i] = newV
	}

	fmt.Printf("%.4f", nums[0]/nums[1])
}
