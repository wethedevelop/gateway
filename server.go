package main

import "github.com/gin-gonic/gin"

// 路由绑定在这里
func Route() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
