package auth

import (
	"net/http"
	"oblivion/pkg/model"
	"oblivion/pkg/session"

	"github.com/gin-gonic/gin"
)

func logoutGet(c *gin.Context) {
	session.Default(c, "session", &model.Session_model{}).Delete(c)

	c.Redirect(http.StatusSeeOther, "/auth/login")
}
