package main

import "fmt"

//functions have their type declared after there name in the parameter space
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

//functions can delcare returned type as well
func f2() (anything string) {
	anything = "hello"
	return
}

//returning multiple values... see the call below
func f() (int, int) {
	return 5, 6
}

//where we take the args with that keyword but make it flexible the ...int) int to take n ahy number of ints
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

//closure exmaple
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

//recursion exmaple
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	fmt.Println("Check this out: ", f2())

	x, y := f()
	fmt.Println("Checkout x: ", x)
	fmt.Println("Checkout y: ", y)

	fmt.Println(add(1, 2, 3))
	add2 := func(x, y int) int {
		return x + y
	}
	fmt.Println("'Add2' gives: ", add2(1, 1))

	//closure
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
	fmt.Println(factorial(4))

	first := func() {
		fmt.Println("1st")
	}
	second := func() {
		fmt.Println("2nd")
	}

	//defer will cause the second function to execute at the end of the current parent function. So first will execute then second.
	defer second()
	first()

}
