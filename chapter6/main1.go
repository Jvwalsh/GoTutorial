package main

import "fmt"

// func main() {
// 	var x [5]float64
// 	x[0] = 98
// 	x[1] = 93
// 	x[2] = 77
// 	x[3] = 82
// 	x[4] = 83

// 	var total float64 = 0
// 	for i := 0; i < len(x); i++ {
// 		total += x[i]
// 	}
// 	fmt.Println(total / float64(len(x)))
// }
func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))

	b := make([]float64, 5, 100)
	fmt.Println("The length of B is: ", len(b))

	arr := [5]float64{1, 2, 3, 4, 5}

	newArr := arr[1:3]
	fmt.Println("check out: ", newArr)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

}

//SLICES
//flexible array of float64 types named x
// var x []float64
//length of the starting slice
// x := make([]float64, 5)
//max length of 10
// x := make([]float64, 5, 10)

//another format is high/low slicing
// arr := [5]float64{1,2,3,4,5}
// arr := [5]float64{1,2,3,4,5}
// x := arr[0:5] -> this will slice that from [0,5) .
// x := arr[1:3] -> this will slice that from [1,3) soooooo 2,3,
// x := arr[:3]
// x := arr[:]
// x := arr[3:]
