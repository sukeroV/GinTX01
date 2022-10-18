package main

import (
	"errors"
	"fmt"
	"math"
)

// 自定义
type areaError struct {
	// 错误信息
	err string
	// 错误有关的长度
	length int
	// 错误有关的宽度
	width int
}

func (areaerror *areaError) Error() string {
	return fmt.Sprintf("error is :%v ----,长度是 %d，宽是 %d,正确范围是[0,100]", areaerror.err, areaerror.length, areaerror.width)
}
func rectangleArea(a, b int) (int, error) {
	if (a < 0 || a > 100) || (b < 0 || b > 100) {
		return 0, &areaError{"length or width is negative", a, b}
	}
	return a * b, nil
}

//-----------------------------

func main() {
	//一般用法
	//aError(2000)
	//---------

	//自定义
	v, err := rectangleArea(2, 101)
	if err != nil {
		print(err.Error())
		return
	}
	fmt.Printf("%v", v)
	//----------
}

// 一般用法
func aError(v float32) (errRE error) {
	defer func() {
		if e := recover(); e != nil && v > 1000 {
			fmt.Printf("%v", e)
		}
	}()
	if v < 0 {
		panic("不能小于0")
	}
	if v > 100 && v <= 1000 {
		s := errors.New("不能大于100")
		fmt.Printf("%v", s)
		return
	}
	if v > 1000 {
		panic("不能大于1000")
	}

	fmt.Printf("%v", 2*v*v*math.Pi)
	return nil
}

//-------------