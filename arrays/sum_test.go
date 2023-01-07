package arrays

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Parallel()

	t.Run("collection of nay size", func(t *testing.T) {
		t.Parallel()

		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
