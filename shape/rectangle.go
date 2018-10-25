package shape

type Rectangle struct {
	Width  float64
	Height float64
}

func (rectangle Rectangle) CalculateArea() float64 {
    return rectangle.Width * rectangle.Height
}

func (rectangle Rectangle) CalculatePerimeter() float64 {
    return rectangle.Width*2 + rectangle.Height*2
}
