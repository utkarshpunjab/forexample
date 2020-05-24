package main

import (
	"fmt"
)

// CustomErr is a structure
type CustomErr struct {
	name string
}

// This is
func (c CustomErr) Error() string {
	return fmt.Sprintf("This is error message coming from customErr's Error method %v", c.name)
}

//Foo function takes error parameter and prints it
func Foo(e error) {
	fmt.Println("Hello from foo", e)
}

func main() {
	cs := CustomErr{
		name: "James Bond",
	}
	Foo(cs)
}
