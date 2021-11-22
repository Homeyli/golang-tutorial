package main

import "fmt"

func main() {

	nums := []int{50, 92, 85, 20}

	// print len

	fmt.Printf("len of array is %d \n", len(nums))

	for i, v := range nums {

		fmt.Printf("value of index %d is %d...\n", i, v)
	}

	// ----------------------------------------------

	person := make(map[interface{}]interface{})

	person["name"] = 55
	person[3] = "wow!"

	for i, v := range person {
		fmt.Printf("value map of index %v is %v...\n", i, v)
	}
}
