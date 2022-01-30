package main

import "fmt"

func main() {

	grade := []int{90, 50, 60, 40, 35, 60, 80, 72, 60}
	studentOne := Student{"SS001", "Precious", "Lois", 28,
		[]int{70, 50, 60, 40, 85, 90, 98, 78, 64}, "Phoenix", 001}
	studentOne.setId("SS002")
	id := studentOne.getId()
	fmt.Println(id)
	studentOne.setFirstName("Kindness")
	fn := studentOne.getFirstName()
	fmt.Println(fn)
	studentOne.setLastName("Pearl")
	ln := studentOne.getLastName()
	fmt.Println(ln)
	age := studentOne.setAndGetAge(23)
	fmt.Println(age)
	gd := studentOne.getGrade(grade)
	fmt.Println(gd)
	studentOne.setCohortName("Maven")
	cn_ := studentOne.getCohortName()
	fmt.Println(cn_)
	studentOne.setCohortNumber(6)
	c_num := studentOne.getCohortNumber()
	fmt.Println(c_num)

	//fmt.Println(studentOne)

}
