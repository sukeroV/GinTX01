package main

import (
	"fmt"
	"reflect"
)

type MosterInreflectTS struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float64 `json:"score"`
	Sex   string  `json:"sex"`
}

func (mo *MosterInreflectTS) SetName(n string) {
	mo.Name = n
	fmt.Printf("%v\n", mo.Name)
}
func (mo MosterInreflectTS) SetAge(n int) {
	mo.Age = n
	fmt.Printf("%v\n", mo.Age)
}
func (mo MosterInreflectTS) PrintAge() {
	fmt.Printf("%v\n", mo.Age)
}

func main() {
	mosterInreflectTS := MosterInreflectTS{
		Name:  "钟山",
		Age:   80,
		Score: 100.0,
		Sex:   "女",
	}
	//mosterInreflectTS2 := MosterInreflectTS{
	//	Name:  "李四",
	//	Age:   80,
	//	Score: 100.0,
	//	Sex:   "女",
	//}

	//mosterInreflectTS.SetName("000")
	//println(mosterInreflectTS.Name)
	//println(mosterInreflectTS2.Name)
	//
	//mosterInreflectTS.SetAge(10)
	//println(mosterInreflectTS.Age)
	//println(mosterInreflectTS2.Age)

	typeOf := reflect.TypeOf(mosterInreflectTS)
	valueof := reflect.ValueOf(mosterInreflectTS)

	//fmt.Printf("typeof %v////////---TYPENumField is %v\n", typeOf, typeOf.NumField())
	//fmt.Printf("valueof %v///////---VALNumField is %v\n", valueof, valueof.NumField())
	//for i := 0; i < typeOf.NumField(); i++ {
	//	fmt.Printf("TYp///Field%v\n", typeOf.Field(i))
	//	fmt.Printf("valueof///Field%v\n", valueof.Field(i))
	//}
	fmt.Printf("typeOf.NumMethod() --->%v\n", typeOf.NumMethod())
	fmt.Printf("valueof.NumMethod() --->%v\n", valueof.NumMethod())
	valueof.Method(0).Call(nil)
	var values []reflect.Value
	prams := append(values, reflect.ValueOf(100))
	valueof.Method(1).Call(prams)
}
