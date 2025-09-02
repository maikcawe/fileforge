package main

import {
	"fmt"
	"os"
}

func main() {

	fmt.Println("Fileforge is alive! 🚀")

	var path string

	if len(os.Args) == 1 {
		path = "."
	} else {
		path = os.Args[1]
	}

	fmt.Println("Path:", path)

	entries, err := os.ReadDir(path)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}