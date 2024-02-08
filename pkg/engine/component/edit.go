package component

import (
	"encoding/json"
	"net/http"
	"oblivion/pkg/crud"
	"oblivion/pkg/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func editComponentGet(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("user") == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
		return
	}

	user := user.User{}
	if err := json.Unmarshal(session.Get("user").([]byte), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	element, err := crud.GetElement(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.HTML(200, "edit.html", gin.H{
		"element":         element,
		"IsAuthenticated": true,
	})
}

func editComponentPost(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("user") == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
		return
	}

	userinfo := user.User{}
	if err := json.Unmarshal(session.Get("user").([]byte), &userinfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	element, err := crud.GetElement(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	elem := user.Element{
		Id:        c.PostForm("id"),
		Name:      c.PostForm("name"),
		Content:   c.PostForm("content"),
		Remind:    element.Remind,
		Frequency: element.Frequency,
	}
	

	err = crud.UpdateElement(elem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.Redirect(http.StatusSeeOther, "/component/list")
}