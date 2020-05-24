// Package word provide custom functions
package word

import (
	"strings"
)

// UseCount no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count function returns the number of words in the string passed to this function
func Count(s string) int {
	// write the code for this func
	xi := strings.Split(s, " ")
	return len(xi)
}
