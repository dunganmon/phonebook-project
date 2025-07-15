package main

import (
	"TODO/configs"
	"TODO/databases"
	"TODO/models"
	"TODO/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	databases.ConnectDb()
	err := databases.DB.AutoMigrate(models.User{})
	if err != nil {
		return
	}
	router := gin.Default()

	configs.Cors(router)
	routers.Init(router)

	router.Run(":8080")
}
