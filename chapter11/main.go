package main

import "fmt"
import "gobook/chapter11/math"

func main() {
	xs := []float64{1, 2, 3, 4, 600}

	avg := math.Average(xs)
	fmt.Println(avg)

	low := math.Min(xs)
	fmt.Println(low)

	high := math.Max(xs)
	fmt.Println(high)

	varNum2 := math2.VarNum2(xs)
}
