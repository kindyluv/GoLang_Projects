package main

import (
	"fmt"
)

type Class struct {
	Name  string
	Age   int
	Email string
}

func (c *Class) Anything(anything interface{}) interface{} {
	return anything
}

func main() {

	mapping := make(map[string]interface{})
	mapping["Name"] = "Precious"
	mapping["Age"] = 29
	mapping["Email"] = "precious@gmail.com"

	for k, v := range mapping {
		fmt.Printf("Key: %s and Value: %v\n", k, v)
	}

	result := Class{
		Name:  "Lois",
		Age:   30,
		Email: "lois@example.com",
	}

	result.Anything(mapping)
	fmt.Println(result)
	fmt.Println(mapping)
}
