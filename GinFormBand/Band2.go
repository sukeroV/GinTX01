package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// 127.0.0.1:8080/register?username=he&phone=1234567890&password=121212
func main() {
	r := gin.Default()
	r.POST("/register", func(context *gin.Context) {
		var register Reguster
		if err := context.ShouldBind(&register); err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Printf("%v", register.Phone)
		context.Writer.Write([]byte(register.UserName + "  " + register.Phone))
	})
	r.Run()
}

type Reguster struct {
	UserName string `form:"username"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}
