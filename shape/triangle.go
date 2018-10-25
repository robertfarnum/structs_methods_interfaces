package shape

// Triangle has three points of fun
type Triangle struct {
	Base   float64
	Height float64
}

// CalculateArea is a function to calculate ... hmm let's see Area
func (triangle Triangle) CalculateArea() float64 {
    return triangle.Height * triangle.Base / 2.0
}
