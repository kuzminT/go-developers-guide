type Container struct {
	sync.Mutex
	Mp map[int][]int
}

func readFromChanAndWriteResult(x1 int, x2 int, fn func(int) int, container *Container, i int, wg sync.WaitGroup) {
	container.Mutex.Lock()
	defer container.Mutex.Unlock()
    container.Mp[i] = []int{fn(x1), fn(x2)}
	wg.Done()
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {

	go func() {
        defer close(out)
		mp := make(map[int][]int)
		container := Container{Mp: mp}
		var wg sync.WaitGroup
		wg.Add(n)
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2

            go func() {
                readFromChanAndWriteResult(x1, x2, fn, &container, i, wg)
                      //readFromChanAndWriteResult(x2, fn, &container, i, wg)}()
		}
		wg.Wait()
        go func() {
            for i := 0; i < n; i++ {
                out <- mp[i][0] + mp[i][1]
            }
        }()
	}()
}