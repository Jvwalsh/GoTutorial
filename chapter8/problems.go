package main

import "fmt"

// How do you get the memory address of a variable?
//Access the address using the & operator before the variable

// How do you assign a value to a pointer?
//prefix the new value with *

// How do you create a new pointer?
//using the new keyword. eg.,
// xPtr := new(int)
// AddOne(xPtr)
// fmt.Println(*xPtr) // x is 1

// What is the value of x after running this program:

// func square(x *float64) {
//   *x = *x * *x
// }
// func main() {
//   x := 1.5
//   square(&x)
// }
//the value after running this program is: 1.5*1.5 2.25 (confirmed below)

// Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).

func square(x *float64) {
	*x = *x * *x
}
func swap(a *int, b *int) {

	holding := *a
	*a = *b
	*b = holding
}

func main() {
	x := 1.5
	square(&x)
	fmt.Println("x is : ", x)
	a := 10
	b := 24
	fmt.Println("a is: ", a)
	fmt.Println("b is: ", b)
	swap(&a, &b)
	fmt.Println("after swap")
	fmt.Println("a is: ", a)
	fmt.Println("b is: ", b)
}
