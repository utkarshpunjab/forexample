package main

import "testing"

func TestMysum(t *testing.T) {
	type test struct {
		nums []int
		ans  int
	}

	tests := []test{
		test{[]int{22, 20}, 42},
		test{[]int{10, 11}, 21},
		test{[]int{2, 7}, 9},
	}
	for _, v := range tests {
		if v.ans != mySum(v.nums...) {
			t.Error("Expected result was", v.ans, "but actual result is", mySum(v.nums...))
		}
	}
}
