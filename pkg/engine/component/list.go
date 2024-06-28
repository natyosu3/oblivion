package component

import (
	"net/http"
	"oblivion/pkg/crud"
	"oblivion/pkg/model"
	"oblivion/pkg/session"

	"github.com/gin-gonic/gin"
)

func listComponentGet(c *gin.Context) {
	data, ok := session.Default(c, "session", &model.Session_model{}).Get(c).(*model.Session_model)

	if data == nil || !ok {
		c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
		return
	}

	elements, err := crud.GetListElement(data.User.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.HTML(200, "list.html", gin.H{
		"elements":        elements,
		"IsAuthenticated": true,
	})
}
