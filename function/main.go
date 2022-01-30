package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func add(x, y int) (z1, z2 int) {
	defer fmt.Println("Hello, world!")

	z1 = x + y
	z2 = x - y
	fmt.Println("Before return")
	return
}

func main() {
	//x := 5
	//zero(&x)
	//fmt.Println(x) // x is 0

	ans1, ans2 := add(23, 5)
	fmt.Printf("X + Y = %d \nX - Y = %d", ans1, ans2)
}
