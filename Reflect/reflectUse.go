package main

import (
	"fmt"
	"reflect"
)

type Moster struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float64 `json:"score"`
	Sex   string  `json:"sex"`
}

func (s Moster) Print() {
	fmt.Printf("--------stsrt---------\n")
	fmt.Println(s)
	fmt.Println("----------enf-----------")
}
func (s Moster) GetSum(n1, n2 int) int {
	return n1 + n2
}
func (s Moster) Set(name string, age int, score float64, sex string) {
	s.Name = name
	s.Age = age
	s.Sex = sex
	s.Score = score
}
func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()

	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	num := val.NumMethod()
	fmt.Printf("num have %v\n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("field %v,tag=%v    %v \n", i, val.Field(i), typ.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("%v,tag为=%v、n", i, tagVal)
		}
	}

	numOgMethod := val.NumMethod()
	fmt.Printf("%v \n", numOgMethod)
	val.Method(1).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Printf("res=%v\n", res[0].Int())

}
func main() {

	var a Moster = Moster{
		Name:  "饼干",
		Age:   10,
		Score: 38.5,
	}
	TestStruct(a)
}
