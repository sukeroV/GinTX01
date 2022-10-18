package main

import "fmt"

var Fornum int = 8000

func main() {
	numChan := make(chan int, 2000000000)
	resChan := make(chan int, Fornum)
	exitChan := make(chan bool, 1)

	go numChanFC(numChan)
	for i := 0; i < Fornum; i++ {
		go resChanFC(numChan, resChan, exitChan, i)
	}

	if <-exitChan {
		go PrintresChan(resChan, exitChan)
	}

	for {
		v := <-exitChan
		if v == false {
			break
		}
	}
}

func numChanFC(numChan chan int) {
	for i := 1; i <= 20000; i++ {
		numChan <- i
	}
	close(numChan)
}
func resChanFC(numChan chan int, resChan chan int, exitChan chan bool, int2 int) {
	addV := 0
	for {
		v, ok := <-numChan
		if !ok {
			break
		}
		addV += v
		fmt.Printf("%v+%v  ", int2, addV)
	}
	resChan <- addV
	if len(resChan) == Fornum {
		close(resChan)
		exitChan <- true
	}
}

func PrintresChan(resChan chan int, exitChan chan bool) {
	iiiiiii := 0
	iiiiibb := 0
	for {
		v, ok := <-resChan
		if !ok {
			break
		}
		if iiiiibb%2 == 0 {
			println()
		}
		fmt.Printf("%v/// ", v)
		iiiiibb++
		iiiiiii += v
	}
	exitChan <- false
	close(exitChan)
	fmt.Printf("\n%v", iiiiiii)
}
