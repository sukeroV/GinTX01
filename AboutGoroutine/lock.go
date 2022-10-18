package main

//import (
//	"fmt"
//	"runtime"
//	"sync"
//)
//
//func main() {
//
//}
//
///*
//加锁
//计算1-200阶乘
//思路：
//1、函数计算放到map
//2、启动协程，放到map
//3、map全局
//*/
////全局锁
//var (
//	myMap = make(map[int]uint64, 10)
//	lock  sync.Mutex //全局互斥锁 Mutex
//)
//
//func Jiechen(n int) {
//
//	var res uint64 = 1
//	for i := 1; i <= n; i++ {
//		res *= uint64(i)
//	}
//	lock.Lock() //加锁
//	//放到map
//	myMap[n] = res
//	lock.Unlock() //解锁
//
//}
//
////调用方式
//func JiechenDY() {
//	for i := 1; i <= 20; i++ {
//		go Jiechen(i)
//	}
//	//遍历结果
//	lock.Lock()
//	for i, v := range myMap {
//
//		fmt.Printf("%v,%v\n", i, v)
//	}
//	lock.Unlock()
//	fmt.Printf("%v\n", myMap[10])
//}
//
///*
//创建catChan，方式个结构体变量，演示写入何读取
//*/
//type xcCat struct {
//	Name string
//}
//
//func catChanFC() {
//	var catChan chan interface{}
//	catChan = make(chan interface{}, 10)
//	cat1 := xcCat{Name: "he"}
//	cat2 := xcCat{Name: "sh"}
//	catChan <- cat1
//	catChan <- cat2
//	//取出、
//	cat11 := <-catChan
//	cat22 := <-catChan
//	//fmt.Printf("%v,%v", cat11.Name, cat22) //需要类型断言，不断的话cat11与22默认还是空接口
//	fmt.Printf("%v,%v", cat11.(xcCat).Name, cat22)
//
//}
//
///*
//演示管道的使用
//创建一个可以存放3个int类型的管道
//理解chan：
//
//	chan的
//*/
//func JichuChan() {
//	var intChan chan int
//	intChan = make(chan int, 3)
//
//	fmt.Printf("%v\n", intChan) //0xc00008c000 chan本质是地址
//
//	//3、向管道写入数据
//	intChan <- 10
//	num := 200
//	intChan <- num
//
//	//写入时不能超容量，但是可以取出来再放
//	//管道长度和cap
//	fmt.Printf("len= %v,cap= %v\n", len(intChan), cap(intChan))
//}
//
///*
//CPU核心数
//*/
//func cpuNum() {
//	cpuUnm := runtime.NumCPU()
//	fmt.Printf("cpuNum=%v\n", cpuUnm)
//	runtime.GOMAXPROCS(cpuUnm - 1)
//	fmt.Printf("OK")
//}
