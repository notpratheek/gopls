package main

import (
	"fmt"
	"os"
	"strings"
)

// No loop involved, directly use the Join method of
// string package. Join each cmd line arg with a space,
// and Print it
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
