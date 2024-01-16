package engine

import (
	"github.com/gin-gonic/gin"
	"oblivion/pkg/engine/auth"
	"oblivion/pkg/engine/top"
)

func Engine(r *gin.Engine) *gin.Engine  {
	r.LoadHTMLGlob("web/templates/*/*.html")

	topGroup := r.Group("/")
	{
		topGroup.GET("/", top.Index())
	}

	authGroup := r.Group("/auth")
	{
		authGroup.GET("/login", auth.LoginGet())
		authGroup.POST("/login", auth.LoginPost())

		authGroup.GET("/register", auth.RegisterGet())
		authGroup.POST("/register", auth.RegisterPost())
	}

	return r
}