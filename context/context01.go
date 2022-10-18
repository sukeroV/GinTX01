package main

import (
	"fmt"
	"time"
)

var notify bool

func f(b chan bool) {

	i := 0
	for {
		fmt.Printf("%v\n", i)
		i++
		time.Sleep(1 * time.Second)
		if notify {
			break
		}
	}
	defer func() { b <- true }()
}

func main() {
	exitChan := make(chan bool, 1)
	go f(exitChan)

	for {
		
		if <-exitChan {
			break
		}
	}
}
