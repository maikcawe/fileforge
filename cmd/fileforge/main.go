package main

import "fmt"
import "os"

func main() {

	fmt.Println("Fileforge is alive! 🚀")

	var path string

	if len(os.Args) == 1 {
		path = "."
	} else {
		path = os.Args[1]
	}

	fmt.Println("Path:", path)
}