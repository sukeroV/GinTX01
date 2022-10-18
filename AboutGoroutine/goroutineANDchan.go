package main

//
//import "fmt"
//
///*
//协程与管道的应用
//与lock.go 中的lock对比（不推荐使用lock，）
//*/
//func main() {
//	intChan := make(chan int, 50)
//	exitChan := make(chan bool, 1)
//	//fmt.Printf("%v\n", intChan)
//	go WData(intChan)
//	println("14")
//	go RData(intChan, exitChan)
//	//time.Sleep(10 * time.Second)
//	for {
//		_, ok := <-exitChan
//		if !ok {
//			break
//		}
//	}
//}
//
//func WData(intChan chan int) {
//	for i := 0; i < 50; i++ {
//		intChan <- i
//		fmt.Printf("W:%v//  ", i)
//	}
//	fmt.Printf("len: %v\n", len(intChan))
//	close(intChan)
//}
//func RData(intChan chan int, exitChan chan bool) {
//	for {
//		v, ok := <-intChan
//		if !ok {
//			break
//		}
//		fmt.Printf("v:%v// ", v)
//	}
//	//任务完成
//	exitChan <- true
//	close(exitChan)
//}
