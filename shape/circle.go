package shape

import "math"

type Circle struct {
	Radius float64
}

func (circle Circle) CalculateArea() float64 {
    return math.Pi * circle.Radius
}

func (circle Circle) CalculatePerimeter() float64 {
    return 2 * math.Pi * circle.Radius
}
