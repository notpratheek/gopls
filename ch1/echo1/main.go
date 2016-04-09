package main

import (
	"fmt"
	"os"
)

// loop over cmd line args, and print them
// each being seperated by a space
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
