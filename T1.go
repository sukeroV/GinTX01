package main

import "github.com/gin-gonic/gin"

var adminUsers = gin.H{
	"shenz":     gin.H{"email": "sgeb@123.com", "phone": "142341977555"},
	"mashibing": gin.H{"email": "mashibing@123.com", "phone": "142341977555"},
	"lianp":     gin.H{"email": "lianp@123.com", "phone": "142341977555"},
	"glang":     gin.H{"email": "glang@123.com", "phone": "142341977555"},
}

func main() {
	engine := gin.Default()

	group := engine.Group("/admin", gin.BasicAuth(gin.Accounts{
		"shenz":     "123",
		"mashibing": "123",
		"lianp":     "123",
		"glang":     "123",
	}))

	group.GET("/", groupHeadler)

	engine.Run(":9999")
}

func groupHeadler(c *gin.Context) {
	s := c.MustGet(gin.AuthUserKey).(string)
	print("sssssssssssssssssssssssssssssssssssssssssssssssssssssss is-", s)
	a, ok := adminUsers[s]
	if ok {
		c.JSON(200, a)
	} else {
		c.JSON(500, "no")
	}
}
