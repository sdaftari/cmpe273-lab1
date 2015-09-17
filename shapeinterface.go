package shapeinterface

import "math"
import "fmt"

//Structs
type Circle struct {
	X, Y, R float64
}

type Rectangle struct {
	X1, Y1, X2, Y2 float64
}

//Interface
type Shape interface {
	area() float64
	perimeter() float64
}

//Function Implementation
func (c *Circle) area() float64 {
	return math.Pi * c.R * c.R
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.R
}

func (r *Rectangle) area() float64 {
	l := distance(r.X1, r.Y1, r.X1, r.Y2)
	w := distance(r.X1, r.Y1, r.X2, r.Y1)

	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.X1, r.Y1, r.X1, r.Y2)
	w := distance(r.X1, r.Y1, r.X2, r.Y1)

	return 2 * (l + w)
}

func distance(x1, y1, x2, y2 float64) float64 {
	d1 := x2 - x1
	d2 := y2 - y1

	return math.Sqrt(d1*d1 + d2*d2)
}

func TotalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}

	return area
}

func TotalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}

	return perimeter
}