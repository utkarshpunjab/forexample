package word

import (
	"fmt"
	"testing"

	"github.com/utkarshpunjab/forexample/code_samples/testing/ex02/quote"
)

func TestCount(t *testing.T) {
	num := Count("Hello World")
	if num != 2 {
		t.Error("Expected ", 2, "Got ", num)
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("one two two three three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("Expected ", 1, "Got ", v)
			}
		case "two":
			if v != 2 {
				t.Error("Expected ", 2, "Got ", v)
			}
		case "three":
			if v != 3 {
				t.Error("Expected ", 3, "Got ", v)
			}
		}
	}
}

func ExampleCount() {
	fmt.Println(Count("Hello World"))
	// Output
	// 2
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
