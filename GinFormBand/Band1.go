package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()

	engine.GET("hello", func(context *gin.Context) {
		fmt.Printf(".%v", context.FullPath())
		var student Student
		err := context.ShouldBindQuery(&student)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Printf("\n%v\n", student.Name)
		fmt.Printf("\n%v\n", student.Class)
		context.Writer.Write([]byte("hello," + student.Name))

	})
	engine.Run()
}

type Student struct {
	Name  string `form:"name"`
	Class string `form:"class"`
}
