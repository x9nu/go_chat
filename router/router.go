package router

import (
	"go_chat/middleware"
	"go_chat/services"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/login", services.Login)        // user login
	r.POST("/send/code", services.SendCode) // send verification code

	auth := r.Group("/u", middleware.AuthCheck())
	{
		auth.GET("/user/detail", services.UserDetail)
	}

	return r
}
