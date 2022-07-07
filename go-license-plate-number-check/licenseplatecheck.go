package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if _, err := strconv.Atoi(os.Args[1]); err == nil {
		fmt.Printf("%q looks like a number.\n", os.Args[1])
		os.Exit(42)
	} else {
		fmt.Printf("%q is not a number.\n", os.Args[1])
	}
}
