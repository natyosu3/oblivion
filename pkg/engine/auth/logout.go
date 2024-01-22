package auth

import (
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func logoutGet() gin.HandlerFunc {
	return func (c *gin.Context)  {
		session := sessions.Default(c)

		session.Clear()
		session.Save()

		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}