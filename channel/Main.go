package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		c <- 1
	}()

	val := <-c
	fmt.Println("Val: ", val)
	fmt.Println("C channel: ", c)

	go func() {
		c <- 2
	}()

	vaw := <-c
	fmt.Println("Val2: ", vaw)
	fmt.Println("C channel2: ", c)
}
