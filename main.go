package main

import (
	"fmt"
	"net/http"
	"user-server/db"
	"user-server/user"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	router := gin.Default()

	router.LoadHTMLGlob("template/*")
	router.GET("/", func(c *gin.Context) {
		users := []user.User{}
		err := db.Get().Select(&users, "SELECT id, name FROM public.user")
		if err != nil {
			fmt.Println(err)
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"users": users,
		})
	})

	router.GET("/user", user.GetList)
	router.GET("/user/:id", user.Get)
	router.POST("/user", user.Create)
	router.PATCH("/user/:id", user.Edit)
	router.DELETE("/user/:id", user.Delete)

	router.Run(":8080")
}
