package shape

import "math"

// Circle  is a round thing
type Circle struct {
	Radius float64
}

//CalculateArea is a function to calculate ... hmm let's see Area
func (circle Circle) CalculateArea() float64 {
    return math.Pi * circle.Radius
}

// CalculatePerimeter is a function that does the geometric math for a permiter
func (circle Circle) CalculatePerimeter() float64 {
    return 2 * math.Pi * circle.Radius
}
