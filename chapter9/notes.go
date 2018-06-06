package main

import (
	"fmt"
	"math"
)

// func circleArea(c Circle) float64 {
// 	return math.Pi * c.r * c.r
// }

// type Circle struct {
// 	x, y, r float64
// }

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// func main() {
// 	//option 1
// 	// var c Circle

// 	//option 2
// 	// c := new(Circle)

// 	//assign values outright
// 	c := Circle{x: 0, y: 0, r: 5}
// 	//or
// 	// c := Circle{0, 0, 5}
// 	fmt.Println(circleArea(c))

// 	fmt.Println("C is: ", c)

// }

// func (c *Circle) area() float64 {
// 	return math.Pi * c.r * c.r
// }

func main() {

	// c := Circle{0, 0, 5}

	// fmt.Println(c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	robot := new(Android)
	robot.Name = "JOhn"
	robot.Person.Name = "JOhn" //same
	robot.Person.Talk()
	robot.Talk() //same

}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person //person here gets no type and se is an embedded type
	Model  string
}

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}
