package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Edit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}
