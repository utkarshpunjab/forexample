package saying

import "testing"

func TestGreet(t *testing.T) {
	s := Greet("Todd Mutt")
	if s != "Welcome to the mansion Todd Mutt" {
		t.Error("Expected ", "Welcome to the mansion Todd Mutt", "Got ", s)
	}
}

func ExampleGreet() {
	Greet("Christina Rose")
	// Output
	// Welcome to the mansion Christina Rose
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Todd Mutt")
	}
}
