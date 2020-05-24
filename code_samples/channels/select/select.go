package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//Send
	go send(even, odd, quit)
	//Receive
	func(e <-chan int, o <-chan int, q <-chan int) {
		for {
			select {
			case v := <-e:
				fmt.Println("This is even number: ", v)
			case v := <-o:
				fmt.Println("This is Odd number: ", v)
			case v := <-q:
				fmt.Println("This is zero number: ", v)
				return
			}
		}
	}(even, odd, quit)
	fmt.Println("about to exit")
}

func send(e chan<- int, o chan<- int, q chan<- int) {
	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
}
