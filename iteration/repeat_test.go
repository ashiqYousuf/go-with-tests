package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {
		got := Repeat("A", 5)
		want := "AAAAA"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("repeat 0 times", func(t *testing.T) {
		got := Repeat("", 0)
		want := ""

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("repeat 0 times", func(t *testing.T) {
		got := Repeat("", 30000000)
		want := ""

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("repeat 0 times", func(t *testing.T) {
		got := Repeat("abc", -10)
		want := ""

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}

// ? Benchmarking in Go is a way to measure the performance of your code
// ? Benchmarking helps you identify which implementation is more efficient in terms of execution time.
// ! go test -bench=.

func BenchmarkRepeat1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("", 500000000)
	}
}

func BenchmarkRepeat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatOptimized("", 500000000)
	}
}
