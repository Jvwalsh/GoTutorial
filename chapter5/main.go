package main

import "fmt"

// func main() {
// 	i := 0
// 	for i < 11 {
// 		fmt.Println(i)
// 		i++
// 	}
// }
// func main() {
// 	for i := 0; i < 11; i++ {
// 		if i%2 == 0 {
// 			fmt.Println("even")
// 		} else {
// 			fmt.Println("odd")
// 		}
// 	}
// }

// func main() {
// 	i := 4
// 	switch i {
// 	case 0:
// 		fmt.Println("Zero")
// 	case 1:
// 		fmt.Println("One")
// 	case 2:
// 		fmt.Println("Two")
// 	case 3:
// 		fmt.Println("Three")
// 	case 4:
// 		fmt.Println("Four")
// 	case 5:
// 		fmt.Println("Five")
// 	default:
// 		fmt.Println("Unknown Number")
// 	}

// }

//Write a program that prints the numbers from 1 to 100. But for multiples of three print "Fizz" instead of the number and for the multiples of five print "Buzz". For numbers which are multiples of both three and five print "FizzBuzz".ˀ

var a = "fizz"
var b = "buzz"
var c = "fizzbuzz"

func main() {
	for i := 0; i < 100; i++ {
		if i%15 == 0 {
			fmt.Println(c)
		} else if i%5 == 0 {
			fmt.Println(b)
		} else if i%3 == 0 {
			fmt.Println(a)
		} else {
			fmt.Println(i)
		}
	}
}
