package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	cr := make(<-chan int)
	cs := make(chan<- int)
	go func() {
		ch <- 32
	}()
	fmt.Printf("ch details are: %T\t%v\n", ch, ch)
	fmt.Printf("cr details are: %T\t%v\n", cr, cr)
	fmt.Printf("cs details are: %T\t%v\n", cs, cs)
	cr = ch
	fmt.Printf("new cr details are: %T\t%v\n", cr, cr)
	fmt.Println(<-cr)
}
