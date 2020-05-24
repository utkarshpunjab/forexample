package main

import "fmt"

func main() {
	fmt.Println("2+3 = ", mySum(2, 3))
	fmt.Println("10+19 = ", mySum(10, 19))
	fmt.Println("12+13 = ", mySum(12, 13))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
