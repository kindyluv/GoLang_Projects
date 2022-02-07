package main

import (
	"fmt"
)

func heavy() {
	//for i := 0; i <= 22; i++ {
	//	time.Sleep(time.Second * 1)
	//	fmt.Println("Heavy")
	//}

	for i := 0; i < 10; i++ {
		fmt.Println("in heavy", i)
	}
}

func superHeavy() {
	//for i := 0; i < 10; i++ {
	//	time.Sleep(time.Second * 2)
	//	fmt.Println("Super Heavy")
	//}
	for {
		fmt.Println("superheavy")
	}
}

func main() {
	//go heavy()
	//go superHeavy()
	//for i := 0; i < 10; i++ {
	//	fmt.Println("in main")
	//}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//var input string
	//_, err := fmt.Scanln(&input)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//time.Sleep(time.Second * 1)
	//fmt.Println(input)
	//wg:= $sync.WaitGroup {}
	//wg.Add(2)
	//wg.Done()

}
