package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Parallel()

	rectangle := RectAngle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Parallel()

	checkArea := func(r *testing.T, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		t.Parallel()

		rectangle := RectAngle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		t.Parallel()

		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
