package function

import "github.com/gin-gonic/gin"

// 这个是为了测试ngrok的使用

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
