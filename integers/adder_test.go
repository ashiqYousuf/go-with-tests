package integers

import (
	"fmt"
	"testing"
)

func assertCorrectMessage(t testing.TB, got, want int) {
	// ? t.Helper() is needed to tell the test suite that this method is a helper
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestAdder(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		got := Add(2, 2)
		want := 4

		assertCorrectMessage(t, got, want)
	})
}
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
}
