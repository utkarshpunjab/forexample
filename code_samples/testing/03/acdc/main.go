// Package acdc is ready to rock
package acdc

// Adder function adds a list of integers and returns the result
func Adder(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
