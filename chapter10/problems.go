package main

import (
	"fmt"
	"time"
)

// How do you specify the direction of a channel type?
//by changing the rearrangement around the " <- " operator
// chan <- msg means send msg to chan
// msg := <- chan means  recve message from chan and store in msg

// Write your own Sleep function using time.After.

func announce() {
	time.Sleep(time.Second * 3)
	fmt.Println("Announcement Here")
}

func main() {
	fmt.Println("begin")
	announce()
}

// What is a buffered channel? How would you create one with a capacity of 20?
//A buffered channel is the same as a regular channel with the added second parameter where the channel is declared (see 20 below)
//It is useful in that it will operate asynchronously insteasd of synchronously until the buffer channel is full. Below could occur 20 times prior to the channel filling
//c := make(chan int, 20)
