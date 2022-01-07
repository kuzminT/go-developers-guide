package main

import "fmt"

func main() {
	pipe := make(chan string, 2)
	go func() {
		pipe <- "Привет, я анонимная горутина"
		pipe <- "Привет, я ешё одна горутина"
		defer close(pipe)
	}()

	fmt.Println("Основная функция")

	fmt.Println(cap(pipe))

	// for i := 0; i < cap(pipe); i++ {
	// 	fmt.Println(<-pipe)
	// }
	for val := range pipe {
		fmt.Println(val)
	}
	// defer close(pipe)
	// ch := make(chan string, 5)

	// go task2(ch, "test")
}

func readFromChan(ch chan string) {
	fmt.Println(<-ch)
}

func task2(ch chan string, s string) {
	defer close(ch)
	s += " "
	for i := 0; i < 4; i++ {
		ch <- s
	}
}
