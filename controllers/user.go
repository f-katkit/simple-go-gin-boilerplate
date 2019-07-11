package controllers

import (
	"fmt"

	"github.com/f-katkit/simple-go-gin-boilerplate/models"
	"github.com/gin-gonic/gin"

	"net/http"
)

type UserController struct{}

var m = new(models.User)

func (pc UserController) Index(c *gin.Context) {
	p, err := m.GetAll()

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}
