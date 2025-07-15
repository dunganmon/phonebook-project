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
	{
		auth.POST("profiles", middlewares.CheckLogin(), controllers.GetUserProfile)
		auth.GET("/contacts", middlewares.CheckLogin(), controllers.GetContacts)
		auth.POST("/contacts", middlewares.CheckLogin(), controllers.AddContact)
		auth.DELETE("/contacts/:id", middlewares.CheckLogin(), middlewares.CheckOwnContact(), controllers.DeleteContact)
		auth.PUT("/contacts/:id", middlewares.CheckLogin(), middlewares.CheckOwnContact(), controllers.EditContact)
	}
}
