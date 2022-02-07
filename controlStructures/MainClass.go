package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	flag := true
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			flag = false
			break
		}
		fmt.Printf("odd numbers: %d", i)
	}
	fmt.Println(flag)

}
