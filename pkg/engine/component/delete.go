package component

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"oblivion/pkg/crud"
	"oblivion/pkg/user"
)

func deleteComponentPost(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("user") == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
		return
	}

	user := user.User{}
	if err := json.Unmarshal(session.Get("user").([]byte), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	elementId := c.PostForm("id")

	err := crud.DeleteElement(user.UserId, elementId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.Redirect(http.StatusMovedPermanently, "/component/list")
}
