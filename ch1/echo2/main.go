package main

import (
	"fmt"
	"os"
)

// slightly better loop, same as echo1, but using
// range instead of len()
func main() {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
