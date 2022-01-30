package main

import "fmt"

func main() {
	x := 5
	var y [6]int
	y[0] = 10
	y[1] = 13
	y[2] = 15
	y[3] = 24
	y[4] = 31
	y[5] = 1
	for k, v := range y {
		fmt.Println(k, v)
		x++
	}

}
