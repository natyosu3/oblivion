package component

import (
	"net/http"
	"oblivion/pkg/crud"
	"oblivion/pkg/model"
	"oblivion/pkg/session"

	"github.com/gin-gonic/gin"
)

func deleteComponentPost(c *gin.Context) {
	data, ok := session.Default(c, "session", &model.Session_model{}).Get(c).(*model.Session_model)

	if data == nil || !ok {
		c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
		return
	}

	elementId := c.PostForm("id")

	err := crud.DeleteElement(data.User.UserId, elementId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.Redirect(http.StatusSeeOther, "/component/list")
}
