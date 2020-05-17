package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var counter int
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Num of goroutines", runtime.NumGoroutine())
	}
	fmt.Println("Num of goroutines", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Count", counter)
}
