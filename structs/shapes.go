package structs

type RectAngle struct {
	Width  float64
	Height float64
}

func Perimeter(rectAngle RectAngle) float64 {
	return 2 * (rectAngle.Width + rectAngle.Height)
}

func Area(rectAngle RectAngle) float64 {
	return (rectAngle.Width * rectAngle.Height)
}
