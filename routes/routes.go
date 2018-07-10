package routes

import "github.com/gin-gonic/gin"

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func init() {
	GetRouter().GET("/pong", Ping)
}
