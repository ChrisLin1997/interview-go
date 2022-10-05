package user

import (
	"fmt"
	"net/http"
	"user-server/db"

	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	users := []User{}
	err := db.Get().Select(&users, "SELECT id, name FROM public.user")

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
