package main

import "fmt"

func main() {

	var x []int = []int{1, 2, 3, 4, 5}
	//y := [3][5]int{{1, 2, 3, 4, 5}, {3, 4, 5, 6, 7}, {6, 8, 9, 7, 0}}
	sum := 1
	for i := 0; i < len(x); i++ {
		sum *= x[i]

	}
	fmt.Println(sum)
	fmt.Println()

}
