package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, repeated, expected string) {
		t.Helper()

		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	}

	t.Run("Repeat 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("Repeat 4 times", func(t *testing.T) {
		repeated := Repeat("b", 4)
		expected := "bbbb"
		assertCorrectMessage(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("c", 10)
	fmt.Println(repeated)
	// Output: cccccccccc
}
