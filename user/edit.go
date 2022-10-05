package user

import (
	"fmt"
	"net/http"
	"user-server/db"

	"github.com/gin-gonic/gin"
)

func Edit(c *gin.Context) {
	id := c.Param("id")
	var body editRequest
	if err := c.ShouldBind(&body); err != nil {
		fmt.Println(err)
		return
	}

	tx, err := db.Get().Beginx()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer tx.Rollback()

	_, err = tx.NamedExec("UPDATE public.user SET name=:name WHERE id=:id", map[string]interface{}{
		"id":   id,
		"name": body.Name,
	})

	if err != nil {
		fmt.Println(2)
		fmt.Println(err)
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"result": true,
	})
}
