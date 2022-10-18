package main

//  go get -u github.com/sukerov/myfunctions  更新方法
import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func refleactTS2(b interface{}) {

	rType := reflect.TypeOf(b)
	fmt.Printf("%v\n", rType)
	rVal := reflect.ValueOf(b)

	iV := rVal.Interface()
	fmt.Printf("%T\n", iV)

	//先断言

	// 这里的 v 是类型的值
	//switch nameiV := iV.(type) {
	//case bool:
	//	fmt.Printf("bool\n")
	//case string:
	//	fmt.Printf("string\n")
	//case int, int64:
	//	fmt.Printf("int or int 64\n")
	//case float32, float64:
	//	fmt.Printf("float32 or 64\n")
	//default:
	//	fmt.Printf("what Unknown is %v\n", reflect.TypeOf(nameiV))
	//}
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("%v\n", stu.Name)
	}
	//myFunctions.SwitchAllTypr(iV)
}

func main() {
	stu := Student{
		Name: "何",
		Age:  14,
	}
	refleactTS2(stu)

}
