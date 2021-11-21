package main

import "fmt"

type Person struct {
	name   string
	family string
	old    int
}

func main() {

	mahdi := Person{name: "mahmoud", family: "homeyli", old: 31}

	//fmt.Println(mahdi)

	// --------------------------------------------------------------

	var me Person
	me.name = "mahdi"
	me.family = "homeyli"
	me.old = 31

	//fmt.Println(me)

	// -------------------------------------------------------------

	var you [2]Person

	you[0] = me
	you[1] = mahdi

	fmt.Println(you)

	// -------------------------------------------------------------

	//fake := [2]Person{me, mahdi}
	fake := [2]Person{{name: "test", family: "homeyli", old: 31}, {name: "mahmoud", family: "homeyli", old: 31}}

	fmt.Println(fake)

}
