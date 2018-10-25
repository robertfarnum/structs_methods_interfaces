package main

import "github.com/robertfarnum/structs_methods_interfaces/geometry"
import "github.com/robertfarnum/structs_methods_interfaces/shape"

import "testing"

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t *testing.T, calculator geometry.PerimeterCalculator, want float64) {
		t.Helper()
		got := calculator.CalculatePerimeter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := shape.Rectangle{Width: 10.0, Height:10.0}
		checkPerimeter(t, rectangle, 40.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := shape.Circle{Radius: 10.0}
		checkPerimeter(t, circle, 62.83185307179586)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, calculator geometry.AreaCalculator, want float64) {
		t.Helper()
		got := calculator.CalculateArea()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := shape.Rectangle{Width:12, Height:6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := shape.Circle{Radius:10.0}
		checkArea(t, circle, 31.41592653589793)
	})

	t.Run("triangles", func(t *testing.T) {
		triangle := shape.Triangle{Base:10.0, Height:10.0}
		checkArea(t, triangle, 50)
	})
}
