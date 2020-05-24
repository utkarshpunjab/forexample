package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type input struct {
		values []int
		ans    float64
	}

	inputs := []input{
		input{[]int{1, 7, 8, 9, 11}, 8},
		input{[]int{2, 12, 10, 17, 18}, 13},
		input{[]int{1, 5, 10, 4, 21}, 6.333333333333333},
	}

	for _, v := range inputs {
		if CenteredAvg(v.values) != v.ans {
			t.Error("Expected ", v.ans, "Got ", CenteredAvg(v.values))
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 10, 15, 17, 41}))
	// Output
	// 14
}
func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{10, 21, 24, 25, 29, 100}
	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}
