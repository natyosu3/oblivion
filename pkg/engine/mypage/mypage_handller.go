package mypage

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func MypageTop() gin.HandlerFunc {
	return func (c *gin.Context)  {
		session := sessions.Default(c)
		username := session.Get("username")

		if username == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
			c.Abort()
		}

		c.HTML(http.StatusOK, "mypage.html", gin.H{"username": username})
	}
}