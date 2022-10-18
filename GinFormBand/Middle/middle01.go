package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	r := gin.Default()
	r.Use(ReguestInfos())
	r.GET("hello", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  context.FullPath(),
		})
	})
	r.Run()
}

func ReguestInfos() gin.HandlerFunc {

	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method
		fmt.Printf("Path: = %v\n", path)
		fmt.Printf("method: =  %v\n", method)
		context.Next()
		status := context.Writer.Status()
		context.Writer.Write([]byte(strconv.Itoa(status)))
	}
}
