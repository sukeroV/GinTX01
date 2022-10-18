package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("./JsonFormat/*")
	r.Static("/img", "./img") //访问静态资源

	r.GET("/hellojson", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"code":    1,
			"message": "ok",
			"data":    context.FullPath(),
		})
	})

	r.GET("/jsonstruct", func(context *gin.Context) {
		resp := Response{
			Code:    1,
			Message: "OK",
			Data:    context.FullPath(),
		}
		context.JSON(200, &resp)
	})

	r.GET("/helloHTML", func(context *gin.Context) {
		context.HTML(200, "JsonOaout.html", gin.H{
			"fullpath": context.FullPath(),
		})
	})

	r.Run()

}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
