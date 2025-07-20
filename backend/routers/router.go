package routers

import (
	"TODO/controllers"
	"TODO/middlewares"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	noAuth := router.Group("/api")
	{
		noAuth.POST("/register", controllers.Register)
		noAuth.POST("/login", controllers.Login)
	}

	auth := router.Group("/api")
	auth.Use(middlewares.JWTAuth())
	{
		auth.POST("profiles", controllers.GetUserProfile)
		auth.GET("/contacts", controllers.GetContacts)
		auth.POST("/contacts", controllers.AddContact)
		auth.DELETE("/contacts/:id", middlewares.CheckOwnContact(), controllers.DeleteContact)
		auth.PUT("/contacts/:id", middlewares.CheckOwnContact(), controllers.EditContact)
	}
}
