package routes

import "github.com/gin-gonic/gin"

var router *gin.Engine

func GetRouter () *gin.Engine {
	if (router == nil) {
		router = gin.Default()
	}
	return router
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping",
	})
}

func init() {
	GetRouter().GET("/ping", Ping)
}

