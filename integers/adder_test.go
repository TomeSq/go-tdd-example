package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Parallel()

	sum := Add(2, 2)
	excepted := 4

	if sum != excepted {
		t.Errorf("excepted '%d' but got '%d'", excepted, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
