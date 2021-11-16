package main

import (
	"fmt"
	"time"
)

func calc(base int, c chan int) {
	time.Sleep(time.Second)

	c <- 2 * base

	time.Sleep(time.Second * 2)

	c <- 3 * base

	time.Sleep(time.Second * 3)

	c <- 4 * base
}

func main() {
	c := make(chan int)

	go calc(2, c)

	a, b := <-c, <-c

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(<-c)
}
