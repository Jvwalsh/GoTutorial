package main

import "fmt"

//this is no good because panic will execute beofre recover which will end thibngs
// func main() {
// 	panic("PANIC")
// 	str := recover()
// 	fmt.Println(str)
// }

//now we certainly run recovery after panic
func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
