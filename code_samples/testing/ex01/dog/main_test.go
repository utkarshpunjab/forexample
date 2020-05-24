package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	x := Years(10)
	if x != 70 {
		t.Error("Expected ", 70, "Got ", x)
	}
}

func TestYearsTwo(t *testing.T) {
	x := YearsTwo(5)
	if x != 35 {
		t.Error("Expected ", 35, "Got ", x)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output
	// 70
}
func ExampleYearsTwo() {
	fmt.Println(YearsTwo(5))
	// Output
	// 35
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(5)
	}
}
