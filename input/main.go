package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your birth year to get your current age: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Print("Enter your birth month: ")
	scanner.Scan()
	month := scanner.Text()
	fmt.Printf("You will be %d years old by the month of %s", 2022-input, month)

}
