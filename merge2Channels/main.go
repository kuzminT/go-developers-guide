package main

import "sync"

// Необходимо написать функцию func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int).

// Описание ее работы:

// n раз сделать следующее

// прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
// вычислить f(x1) + f(x2)
// записать полученное значение в out
// Функция merge2Channels должна быть неблокирующей, сразу возвращая управление.

// Функция fn может работать долгое время, ожидая чего-либо или производя вычисления.

// Формат ввода:

// количество итераций передается через аргумент n.
// целые числа подаются через аргументы-каналы in1 и in2.
// функция для обработки чисел перед сложением передается через аргумент fn.
// Формат вывода:

// канал для вывода результатов передается через аргумент out.

type Container struct {
	sync.Mutex
	Mp map[int][]int
}

func readFromChanAndWriteResult(x int, fn func(int) int, mp sync.Map, i int, wg sync.WaitGroup) {
	mp.Store(i, fn(x))
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	mp := sync.Map{}
	var wg sync.WaitGroup

	go func() {

		wg.Add(n)
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2
			go func(i int) {
				go readFromChanAndWriteResult(x1, fn, mp, i, wg)

				defer wg.Done()
			}(i)

			go func(i int) {
				go readFromChanAndWriteResult(x2, fn, mp, i+1, wg)

				defer wg.Done()
			}(i)
		}

	}()
	wg.Wait()
	for i := 0; i < n; i++ {
		x1, _ := mp.Load(i)
		x2, _ := mp.Load(i + 1)
		out <- x1.(int) + x2.(int)
	}

}
