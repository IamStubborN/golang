package main

import (
	"fmt"
	"math"
)

type (
	// Shapes interface
	Shapes interface {
		Area() float64
		Perimeter() float64
	}
	// Rectangle struct
	Rectangle struct {
		height, length int
	}
	// Circle struct
	Circle struct {
		radius float64
	}
	// Triangle struct
	Triangle struct {
		sideA float64
		sideB float64
		sideC float64
	}
)

func main() {
	shapes := []Shapes{
		&Rectangle{3, 4},
		&Circle{5},
		&Triangle{3, 4, 5},
	}

	for _, shape := range shapes {
		printShapes(shape)
	}
}

// Perimeter func
func (r *Rectangle) Perimeter() float64 {
	return float64((r.height + r.length) * 2)
}

// Area func
func (r *Rectangle) Area() float64 {
	return float64(r.height * r.length)
}

func (r *Rectangle) String() string {
	return fmt.Sprintf("height = %v, length = %v. Type = %T", r.height, r.length, r)
}

// Perimeter func
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Area func
func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *Circle) String() string {
	return fmt.Sprintf("radius = %v. Type = %T", c.radius, c)
}

// Perimeter func
func (t *Triangle) Perimeter() float64 {
	return t.sideA + t.sideB + t.sideC
}

// Area func
func (t *Triangle) Area() float64 {
	halfPerim := t.Perimeter() / 2
	res := math.Sqrt(halfPerim * (halfPerim - t.sideA) * (halfPerim - t.sideB) * (halfPerim - t.sideC))
	return res
}

func (t *Triangle) String() string {
	return fmt.Sprintf("a = %v, b = %v, c = %v. Type = %T", t.sideA, t.sideB, t.sideC, t)
}

func printShapes(shape Shapes) {
	fmt.Printf("%v\n", shape)
	fmt.Printf("Perimeter = %.2f\n", shape.Perimeter())
	fmt.Printf("Area = %.2f\n----------\n", shape.Area())
}
