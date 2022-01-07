package main

import (
	"fmt"
)

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	result := make(chan int)
	var sum int
	var arg int

	go func() {
		defer close(result)

		for {
			select {
			case arg = <-arguments:
				sum += arg
			case <-done:
				result <- sum
				return
			}
		}
	}()

	return result
}

func main() {
	arguments := make(chan int)
	done := make(chan struct{})
	result := calculator(arguments, done)
	for i := 0; i < 10; i++ {
		arguments <- 1
	}
	close(done)
	fmt.Println(<-result)
}
