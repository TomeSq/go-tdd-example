package structs

import "math"

type Shape interface {
	Area() float64
}

type RectAngle struct {
	Width  float64
	Height float64
}

func (r RectAngle) Area() float64 {
	return (r.Width * r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectAngle RectAngle) float64 {
	return 2 * (rectAngle.Width + rectAngle.Height)
}
