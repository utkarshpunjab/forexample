package main

import "testing"

func TestMysum(t *testing.T) {
	x := mySum(2, 3)
	if x != 5 {
		t.Error("Expected result was", 5, "but actual result is", x)
	}
}
