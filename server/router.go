package server

import (
	"github.com/f-katkit/simple-go-gin-boilerplate/controllers"
	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		u := v1.Group("/users")
		{
			user := new(controllers.UserController)
			u.GET("", user.Index)
		}
	}
	return r
}
