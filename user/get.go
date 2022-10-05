package user

import (
	"fmt"
	"net/http"
	"user-server/db"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	id := c.Param("id")
	user := User{}
	err := db.Get().Get(&user, "SELECT id, name FROM public.user WHERE id=$1", id)

	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, user)
}
