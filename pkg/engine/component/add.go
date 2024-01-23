package component

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"oblivion/pkg/crud"
	"oblivion/pkg/user"
)

func addComponentGet(c *gin.Context) {
	c.HTML(200, "add.html", gin.H{
		"IsAuthenticated": true,
	})
}

func addComponentPost(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("user") == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
		return
	}

	user := user.User{}
	if err := json.Unmarshal(session.Get("user").([]byte), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	elementName := c.PostForm("name")
	content := c.PostForm("content")

	err := crud.InsertElement(user.UserId, elementName, content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.HTML(200, "add.html", gin.H{
		"message"		 : "success",
		"IsAuthenticated": true,
	})
}
