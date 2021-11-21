package main

import "fmt"

func main() {

	var a [3]int

	//a = 10

	a[0] = 91
	a[1] = 52
	a[2] = 21

	//fmt.Println(a)
	fmt.Println(a[1:])

	names := [2]string{"mahdi", "mahmoud"}
	fmt.Println(names)

	persons := [2][3]string{{"mahdi", "homeily", "ahwaz"}, {"soheil", "elahi", "iran"}}
	fmt.Println(persons, persons[0][2])

	// Slice --------------------------------------------------------------------------

	var b []int = make([]int, 10)
	b[6] = 92
	fmt.Println(b)

}
