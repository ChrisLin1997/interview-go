package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "name",
		"id":   "id",
	})
}
