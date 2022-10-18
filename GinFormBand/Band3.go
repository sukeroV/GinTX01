package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.POST("/person", func(context *gin.Context) {
		var person Person
		if err := context.BindJSON(&person); err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Printf("%v", person.Name)
		context.Writer.Write([]byte(person.Name + "  " + person.Sex + "   " + string(person.Age)))
	})
	r.Run()
}

type Person struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int    `json:"age"`
}
