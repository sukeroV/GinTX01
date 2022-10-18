package main

import (
	"fmt"
	"reflect"
)

func refleactTS1(b interface{}) {
	//通过反射获取传入的type，kind，值
	//1、先获取到 refect.type
	rType := reflect.TypeOf(b) //事实上是一个type类型
	fmt.Printf("%v\n", rType)

	//2、获取到value
	rVal := reflect.ValueOf(b)
	fmt.Printf("%T\n", rVal)
	n2 := 2 + rVal.Int()
	fmt.Printf("%v\n", n2)

	//将rVal 转为 interface
	iV := rVal.Interface()
	//将Interface{}断言
	n3 := iV.(int)
	fmt.Printf("%T\n", n3)
}

func main() {
	var num int = 100
	refleactTS1(num)
}
