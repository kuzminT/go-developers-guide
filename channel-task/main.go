package main

import (
	// пакет используется для проверки выполнения условия задачи, не удаляйте его

	"sync"
)

type Container struct {
	sync.Mutex
	Mp map[int][]int
}

func readFromChanAndWriteResult(x int, fn func(int) int, container *Container, i int, l int) {
	container.Mutex.Lock()
	container.Mp[i][l] = fn(x)
	defer container.Mutex.Unlock()

}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {

	go func() {
		var wg sync.WaitGroup
		mp := make(map[int][]int)

		container := Container{Mp: mp}

		for i := 0; i < n; i++ {
			container.Mp[i] = make([]int, 2)
			x1 := <-in1
			x2 := <-in2
			wg.Add(2)

			go func(i int) {
				readFromChanAndWriteResult(x1, fn, &container, i, 0)
				defer wg.Done()
			}(i)

			go func(i int) {

				readFromChanAndWriteResult(x2, fn, &container, i, 1)
				defer wg.Done()
			}(i)

		}
		wg.Wait()

		for i := 0; i < n; i++ {
			out <- container.Mp[i][0] + container.Mp[i][1]

		}

	}()

}

// go func(ch <-chan int, i int, j int) {

// 	num := <-ch
// 	calculate(num, fn, &container, i, j)
// 	defer wg.Done()
// }(in1, i, 0)
// //}

// go func(ch <-chan int, i int, j int) {
// 	num := <-ch
// 	calculate(num, fn, &container, i, j)
// 	defer wg.Done()
// }(in2, i, 1)
