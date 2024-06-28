package component

import (
	"net/http"
	"oblivion/pkg/crud"
	"oblivion/pkg/model"
	"oblivion/pkg/session"
	"time"

	"github.com/gin-gonic/gin"
)

func addComponentGet(c *gin.Context) {
	c.HTML(200, "add.html", gin.H{
		"IsAuthenticated": true,
	})
}

func addComponentPost(c *gin.Context) {
	data := session.Default(c, "session", &model.Session_model{}).Get(c)
	se_data, ok := data.(*model.Session_model)

	if data == nil || !ok {
		c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
		return
	}

	elementName := c.PostForm("name")
	content := c.PostForm("content")

	// 新規追加の場合なので, 次のリマインド日時を24時間後に設定する
	nextRemindTime := time.Now().Add(24 * time.Hour).Format("2006-01-02 15:04:05")

	err := crud.InsertElement(se_data.User.UserId, elementName, content, nextRemindTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(200, "add.html", gin.H{
		"message":         "success",
		"IsAuthenticated": true,
	})
}
