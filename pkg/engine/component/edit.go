package component

import (
	"net/http"
	"oblivion/pkg/crud"
	"oblivion/pkg/model"
	"oblivion/pkg/session"

	"github.com/gin-gonic/gin"
)

func editComponentGet(c *gin.Context) {
	data := session.Default(c, "session", &model.Session_model{}).Get(c).(*model.Session_model)

	if data == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
		return
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
	data := session.Default(c, "session", &model.Session_model{}).Get(c).(*model.Session_model)

	if data == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
		return
	}

	element, err := crud.GetElement(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	elem := model.Element{
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
