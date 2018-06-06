package main

import "fmt"

// How do you access the 4th element of an array or slice?

// What is the length of a slice created using: make([]int, 3, 9)?

// Given the array:

// x := [6]string{"a","b","c","d","e","f"}
// what would x[2:5] give you?

// Write a program that finds the smallest number in this list:

// x := []int{
//   48,96,86,68,
//   57,82,63,70,
//   37,34,83,27,
//   19,97, 9,17,
// }

func main() {
	fmt.Println("Problems!")
	// How do you access the 4th element of an array or slice?
	x := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(x[3])
	// What is the length of a slice created using: make([]int, 3, 9)?
	var y = make([]int, 3, 9)
	fmt.Println(len(y))
	// Given the array:
	z := [6]string{"a", "b", "c", "d", "e", "f"}
	// what would x[2:5] give you?
	fmt.Println(z[2:5])
	xyz := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	smallest := xyz[0]
	for i := 0; i < len(xyz); i++ {
		if smallest > xyz[i] {
			smallest = xyz[i]
		}
	}
	fmt.Println(smallest)
	smallest2 := 1000000
	//using range...
	for test, value := range xyz {
		if smallest2 > value {
			smallest2 = value
			fmt.Println("Current _ is: ", test, " current value is: ", value)
		}
	}
}
