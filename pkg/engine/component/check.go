package component

import (
	"net/http"
	"oblivion/pkg/crud"

	"oblivion/pkg/utils/general"
	"oblivion/pkg/user"

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
	memorization := c.PostForm("memorization")

	if elementId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "elementId is empty"})
		return
	}

	if memorization == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "memorization is empty"})
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

	// memorizationが正なら次回のリマインド日を更新
	if memorization == "yes" {
		elem := user.Element{
			Id:        elementId,
			Name:      element.Name,
			Content:   element.Content,
			Remind:	   general.MakeNextRemindDate(element.Frequency),
			Frequency: element.Frequency + 1,
		}
		err = crud.UpdateElement(elem)
	} else {
		elem := user.Element{
			Id:        elementId,
			Name:      element.Name,
			Content:   element.Content,
			Remind:	   element.Remind,
			Frequency: element.Frequency,
		}
		err = crud.UpdateElement(elem)
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "element not found"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/mypage")
}
