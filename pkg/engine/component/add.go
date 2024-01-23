package component

import (
	"encoding/json"
	"net/http"
	"oblivion/pkg/crud"
	"oblivion/pkg/user"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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

	// 新規追加の場合なので, 次のリマインド日時を24時間後に設定する
	nextRemindTime := time.Now().Add(24 * time.Hour).Format("2006-01-02 15:04:05")

	err := crud.InsertElement(user.UserId, elementName, content, nextRemindTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(200, "add.html", gin.H{
		"message"		 : "success",
		"IsAuthenticated": true,
	})
}
