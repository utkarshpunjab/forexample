package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	go bar()
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Goroutines after wait finishes", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i <= 10; i++ {
		fmt.Println("This is Foo and iteration num: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i <= 10; i++ {
		fmt.Println("This is Bar and iteration num: ", i)
	}
	wg.Done()
}
