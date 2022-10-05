package user

import (
	"fmt"
	"net/http"
	"user-server/db"

	"github.com/gin-gonic/gin"
	"github.com/t-pwk/go-flakeid"
)

type User struct {
	Id   int64  `json:"id,string" db:"id"`
	Name string `json:"name" db:"name"`
}

type createRequest struct {
	Name string `json:"name"`
}

func Create(c *gin.Context) {
	var body createRequest
	if err := c.ShouldBind(&body); err != nil {
		fmt.Println(err)
		return
	}

	g := flakeid.FlakeID{}

	user := User{Id: int64(g.NextID()), Name: body.Name}

	tx, err := db.Get().Beginx()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer tx.Rollback()

	_, err = db.Get().NamedExec("INSERT INTO public.user (id, name) VALUES (:id, :name)", &user)

	if err != nil {
		fmt.Println(err)
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"id":   user.Id,
		"name": user.Name,
	})
}
