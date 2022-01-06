package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var sum int
	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			break
		}
		sum += n
	}
	w := bufio.NewWriter(os.Stdout)
	w.WriteString(strconv.Itoa(sum) + "\n")
	w.Flush()
}
