// Package dog makes us understand about dogs
package dog

import "fmt"

//Years is an exported function that takes human year as input and returns dog year for those many human years
func Years(hy int) int {
	dy := 10 * hy
	fmt.Printf("Number of dog years for human year %v is %v\n", hy, dy)
	return dy
}
