package areacalc

// This package defines a simple area and perimeter calculation for different shapes using interfaces.

import "fmt"

type Circle struct {
	radius float64
}

type rect struct {
	length float64
	width  float64
}

type square struct {
	side float64
}

type shape interface {
	area() float64
	perimeter() float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}
func (r rect) area() float64 {
	return float64(r.length * r.width)
}
func (r rect) perimeter() float32 {
	return float32(2 * (r.length + r.width))
}

func Main() {
	var c Circle = Circle{radius: 5.0}
	var r rect = rect{length: 4.0, width: 3.0}
	// s := square{side: 2.0}
	fmt.Printf("Circle Area: %.2f, Perimeter: %.2f\n", c.area(), c.perimeter())
	fmt.Printf("Rectangle Area: %.2f, Perimeter: %.2f\n", r.area(), r.perimeter())
	// var shapes []shape = []shape{c, r}
	// for _, sh := range shapes {
	// 	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", sh.area(), sh.perimeter())
	// }
}
