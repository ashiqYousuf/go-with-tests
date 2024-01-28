package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of a collection", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9}, []int{})
	want := []int{3, 9, 0}

	// ? Go does not let you use equality operators with slices
	// ? reflect.DeepEqual (not type safe)  which is useful for seeing if any two variables are the same
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("testSumAllTails with non empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("testSumAllTails with empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9}, []int{1})
		want := []int{0, 9, 0}

		checkSums(t, got, want)
	})
}
