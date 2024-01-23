package mypage

import (
	// "log/slog"
	"net/http"

	"oblivion/pkg/user"
	"encoding/json"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func MypageTop() gin.HandlerFunc {
	return func (c *gin.Context)  {
		session := sessions.Default(c)

		if session.Get("user") == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
			return
		}

		user := user.User{}
		if err := json.Unmarshal(session.Get("user").([]byte), &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.HTML(http.StatusOK, "mypage.html", gin.H{"user": user})
	}
}