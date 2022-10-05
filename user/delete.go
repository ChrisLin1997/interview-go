package user

import (
	"fmt"
	"net/http"
	"user-server/db"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id := c.Param("id")

	tx, err := db.Get().Beginx()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer tx.Rollback()

	_, err = tx.NamedExec("DELETE FROM public.user WHERE id=:id", map[string]interface{}{
		"id": id,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"result": true,
	})
}
