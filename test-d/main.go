package main

import f "fmt"

func main() {

	i := 0

	for i <= 10 {

		f.Println(i)
		i++
	}

	f.Println("\n")

	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			break
		}
		f.Println(i)
	}

}
