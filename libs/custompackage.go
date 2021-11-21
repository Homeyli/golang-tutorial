package custompackage

import f "fmt"

func SayHello() {

	f.Println("hellllo")
}

func init() {

	f.Println("hello i'm custom package")
}
