// sum is a function which takes a slice of numbers and adds them together. What would its function signature look like in Go?

// Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).

// Write a function with one variadic parameter that finds the greatest number in a list of numbers.

// Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers.

// The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n).

// What are defer, panic and recover? How do you recover from a run-time panic?
package main

import "fmt"

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (anyName uint) {
		anyName = i
		i += 2
		return
	}
}

func fib(x uint) uint {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return fib(x-1) + fib(x-2)
	}
}

func main() {
	// sum is a function which takes a slice of numbers and adds them together. What would its function signature look like in Go?
	// 	sumProb := func(sliceVar []int) int {
	// 		total := 0
	// 		for _, value := range sliceVar {
	// 			total += value
	// 		}
	// 		return total
	// 	}

	// 	num1 := [4]int{1, 2, 3, 4}
	// 	fmt.Println(sumProb(num1))
	// }
	// Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).

	half := func(num int) (int, bool) {
		fmt.Println("hello input is: ", num)
		return num / 2, num%2 == 0

	}
	fmt.Println(half(3))
	fmt.Println(half(500))
	// Write a function with one variadic parameter that finds the greatest number in a list of numbers.
	greatest := func(args ...int) int {
		big := 0
		for _, value := range args {
			if value > big {
				big = value
			}
		}
		return big
	}
	fmt.Println(greatest(1, 2, 3, 2000, 5))
	// Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers.

	nextOdd := makeOddGenerator()
	// func makeEvenGenerator() func() uint {
	// 	i := uint(0)
	// 	return func() (ret uint) {
	// 	  ret = i
	// 	  i += 2
	// 	  return
	// 	}
	//   }

	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

	// The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n).
	//
	fmt.Println(fib(20))
}
