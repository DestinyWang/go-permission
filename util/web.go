package util

import "github.com/gin-gonic/gin"

func InitRouter() (router *gin.Engine) {
	router = gin.Default()
	router.GET("/hello", hello)
	return router
}

// test
func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
