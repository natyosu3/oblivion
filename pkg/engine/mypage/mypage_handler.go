package mypage

import (
	"net/http"

	"encoding/json"
	"oblivion/pkg/user"

	"oblivion/pkg/crud"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func MypageTop() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("user") == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
			return
		}

		userinfo := user.User{}
		if err := json.Unmarshal(session.Get("user").([]byte), &userinfo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		// 24時間以内にリマインドがある要素を取得する
		elements, err := crud.GetListElement(userinfo.UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		doElements := []user.Element{}
		for _, element := range elements {
			// リマインド日時が24時間以内の要素を取得する
			remindTime, err := time.ParseInLocation("2006-01-02 15:04:05", element.Remind, time.Local)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// 現在の時間から24時間後の時間を計算
			now := time.Now()
			nowPlus24Hours := now.Add(24 * time.Hour)

			// remindTimeが現在の時間から24時間以内であれば追加
			if remindTime.After(now) && remindTime.Before(nowPlus24Hours) {
				doElements = append(doElements, element)
			}
		}

		c.HTML(http.StatusOK, "mypage.html", gin.H{
			"user":            userinfo,
			"elements":        doElements,
			"IsAuthenticated": true,
		})
	}
}
