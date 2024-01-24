package component

import (
	"net/http"
	"oblivion/pkg/crud"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func cheackComponentPost(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("user") == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
		return
	}
	elementId := c.PostForm("id")

	if elementId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "elementId is empty"})
		return
	}

	element := crud.

	c.JSON(http.StatusOK, gin.H{"status": "ok", "elementId": elementId, "message": "elementId is not empty"})
}