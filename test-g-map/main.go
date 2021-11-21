package main

import "fmt"

func main() {

	var informations map[string]string = make(map[string]string)

	informations["name"] = "mahdi"
	informations["family"] = "homeyli"
	informations["site"] = "homeyli.ir"
	informations["old"] = "31"

	fmt.Println(informations, informations["name"])

	numbers := map[string]int{"old": 31, "brith": 1990}

	fmt.Println(numbers)

}
