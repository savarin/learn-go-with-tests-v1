package shapes

import "math"

var count = 0

type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width  float64
    Height float64
}

type Circle struct {
    Radius float64
}

type Triangle struct {
    Width  float64
    Height float64
}

func Perimeter(rectangle Rectangle) float64 {
    return 2 * (rectangle.Width + rectangle.Height)
}

func (rectangle Rectangle) Area() float64 {
    return rectangle.Width * rectangle.Height
}

func (circle Circle) Area() float64 {
    return math.Pi * circle.Radius * circle.Radius
}

func (triangle Triangle) Area() float64 {
    return 0.5 * triangle.Width * triangle.Height
}
