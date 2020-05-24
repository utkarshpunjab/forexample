// Package saying is used to greet people
package saying

import "fmt"

// Greet welcomes the guest name based on the input string passed to it
func Greet(name string) string {
	return fmt.Sprint("Welcome to the mansion ", name)
}
