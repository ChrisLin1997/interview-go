package main

import (
	"user-server/db"
	"user-server/user"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	gin.ForceConsoleColor()

	router := gin.Default()
	router.GET("/user", user.GetList)
	router.GET("/user/:id", user.Get)
	router.POST("/user", user.Create)
	router.PATCH("/user/:id", user.Edit)
	router.DELETE("/user/:id", user.Delete)

	router.Run(":8080")
}
