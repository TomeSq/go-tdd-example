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
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "RectAngle", shape: RectAngle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()

			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %.2f want %.2f", tt, got, tt.want)
			}
		})
	}
}
