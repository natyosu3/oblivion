package component

import (
	"net/http"
	"oblivion/pkg/crud"

	"oblivion/pkg/utils/general"

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

	element, err := crud.GetElement(elementId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "element not found",
			"err":   err,
		})
		return
	}

	nextday := general.MakeNextRemindDate(element.Frequency)

	err = crud.UpdateElement(elementId, element, nextday)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "element not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "elementId": elementId, "message": "elementId is not empty"})
}
