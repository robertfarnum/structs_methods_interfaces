package shape

// Rectangle is a sqaure like thingy
type Rectangle struct {
	Width  float64
	Height float64
}

//CalculateArea is a function to calculate ... hmm let's see Area
func (rectangle Rectangle) CalculateArea() float64 {
    return rectangle.Width * rectangle.Height
}

// CalculatePerimeter is a function that does the geometric math for a permiter
func (rectangle Rectangle) CalculatePerimeter() float64 {
    return rectangle.Width*2 + rectangle.Height*2
}
