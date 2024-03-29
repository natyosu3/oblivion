package auth

import (
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func logoutGet(c *gin.Context) {
	session := sessions.Default(c)

	session.Clear()
	session.Save()

	c.Redirect(http.StatusSeeOther, "/auth/login")
}