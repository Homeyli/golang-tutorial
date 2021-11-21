package main

import (
	"flag"
	f "fmt"
	"os"
)

func main() {

	version := flag.String("version", "1.0.0", "version of app")
	flag.Parse()

	f.Println("cli app ..")
	f.Printf("%v", os.Args)
	f.Println("\n")

	f.Println(*version)
}
