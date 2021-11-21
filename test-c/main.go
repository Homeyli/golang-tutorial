package main

import f "fmt"

func main() {

	//var isLogin bool = false

	/*	if a, b := 10, 20; a != b {
		f.Println("fuck")
	}*/

	number := 0

	switch {

	case number == 20:
		f.Println("perfect")
	case number < 20 && number > 17:
		f.Println("so good")
	case number < 18 && number > 14:
		f.Println("good")
	case number < 15 && number > 9:
		f.Println("OK !")
	case number < 10 && number > 4:
		f.Println("bad")
	case number < 5:
		f.Println("hey std fuck you")
	}

	/*if isLogin {
		f.Println("is OK")
	} else {
		f.Println("no OK")
	}*/
}
