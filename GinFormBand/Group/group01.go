package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	rGroup := r.Group("user")
	rGroup.POST("/register", func(context *gin.Context) {
		fullPath := "用户注册" + context.FullPath()
		fmt.Printf("%v", fullPath)
		context.Writer.Write([]byte(fullPath))
	})
	rGroup.POST("/login", func(context *gin.Context) {
		fullPath := "用户登录" + context.FullPath()
		fmt.Printf("%v", fullPath)
		context.Writer.Write([]byte(fullPath))
	})
	rGroup.DELETE("/:id", func(context *gin.Context) {
		fullPath := "用户删除" + context.FullPath()
		userID := context.Param("id")
		fmt.Printf("%v", userID)
		context.Writer.Write([]byte(fullPath + "  " + userID))
	})

	r.Run()
}
