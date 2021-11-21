package main

import (
	f "fmt"

	custompackage "github.com/homeyli/golang/libs"
)

func add(brith int, name string) (old int, fullname string) {

	return 2021 - brith, name
}

func sum(nums ...int) {

	f.Printf("%v", nums)
}

func main() {

	old, nickname := add(1990, "mahdi homeyli")

	f.Printf("%s is %d years old", nickname, old)
	sum(5, 6, 9, 7, 99, 2, 3, 7, 8)
	custompackage.SayHello()
}
