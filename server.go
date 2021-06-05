package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wethedevelop/gateway/account"
)

// 路由绑定在这里
func Route() *gin.Engine {
	r := gin.Default()

	r.POST("/api/user/signup", account.Signup)
	return r
}
