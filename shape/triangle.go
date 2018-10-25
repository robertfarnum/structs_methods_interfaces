package shape

type Triangle struct {
	Base   float64
	Height float64
}

func (triangle Triangle) CalculateArea() float64 {
    return triangle.Height * triangle.Base / 2.0
}
