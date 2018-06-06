package main

import (
	"fmt"
	"math"
)

// What's the difference between a method and a function?

//A method is assigned to a reicever and specfiically acts with that reciever using the '.' operator.
//A function is used to convert 0 or many input paramters to zero or more output paramters

// Why would you use an embedded anonymous field instead of a normal named field?
//With an embedded anon field, we can represent inheritance to a specific object

// Add a new method to the Shape interface called perimeter which calculates the perimeter of a shape. Implement the method for Circle and Rectangle

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Circle struct {
	x, y, r float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// type MultiShape struct {
// 	shapes []Shape
// }

// func (m *MultiShape) area() float64 {
// 	var area float64
// 	for _, s := range m.shapes {
// 		area += s.area()
// 	}
// 	return area
// }
// func (m *MultiShape) perimeter() float64 {
// 	var perimeter float64
// 	for _, s := range m.shapes {
// 		area += s.area()
// 	}
// 	return area
// }

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (c *Circle) perimeter() float64 {
	return math.Pi * 2 * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l*2 + w*2
}

func main() {

	c := Circle{x: 0, y: 0, r: 5}
	r := Rectangle{2, 4, 10, 10}

	fmt.Println("Perimeters are: ", r.perimeter(), "  ", c.perimeter())
	fmt.Println("Perimeters are: ", r.area(), "  ", c.area())

}
