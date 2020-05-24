package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	//var mu sync.Mutex
	fmt.Println(counter)
	const gs = 100
	fmt.Printf("Num of CPUs %v\t Num of Goroutines %v\n", runtime.NumCPU(), runtime.NumGoroutine())
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			//mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			//mu.Unlock()
			wg.Done()
		}()
		fmt.Printf("Num of CPUs %v\t Num of Goroutines %v\n", runtime.NumCPU(), runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Count", counter)
	//fmt.Printf("Num of CPUs %v\t Num of Goroutines %v\n", runtime.NumCPU(), runtime.NumGoroutine())
}
