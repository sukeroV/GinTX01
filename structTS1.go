package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 100

	fmt.Printf("%T\n", a)
	fmt.Printf("%T", strconv.Itoa(a))
}
