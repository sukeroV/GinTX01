package main

import (
	"errors"
	"fmt"
	"math"
)

func errA(r float32) (fl float32, e error) {
	if r < 0 {
		panic("半径小于0")
	}

	if r < 5 || r > 50 {
		e = errors.New("xiaoyu dayu")
		return
	}
	return r * r * math.Pi * 2, nil
}

func main() {
	defer func() {
		if errRecov := recover(); errRecov != nil {
			fmt.Printf("致死原因：%v", errRecov)
		}

	}()
	mjA, errA := errA(100.0)
	fmt.Printf("%v\n", mjA)
	fmt.Printf("%v", errA)
}
