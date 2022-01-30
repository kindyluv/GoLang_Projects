package main

import (
	"fmt"
)

type Bucket struct {
	name string
	id   int64
}

func test2(myFunc func()) {
	myFunc()
}

func test3(r string) string {
	return r
}

func world() {
	fmt.Println("Hello world!")
}

func hello() {
	fmt.Println("Beg JEHOVAH")
}

func main() {
	test := func(x int) int {
		return x * -1
	}(8)
	fmt.Println(test)

	test2(hello)
	test2(world)

	x := 7
	y := &x
	fmt.Println(x, y)
	*y = 12
	fmt.Println(x, y)

	toChange := "Hello world!"
	var pointer *string = &toChange
	fmt.Println(pointer, &pointer)

}
