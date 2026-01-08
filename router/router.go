package router

import (
	"laporpak/handler"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.GET("/users", handler.GetUsers)
	r.POST("/users", handler.CreateUser)
	r.Run(":8080")
}
