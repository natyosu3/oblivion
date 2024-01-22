package auth

import (
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"oblivion/pkg/crud"
	"oblivion/pkg/utils/crypto"
	"oblivion/pkg/user"
	"encoding/json"
	"log/slog"
)

func loginGet (c *gin.Context) {
	session := sessions.Default(c)

	if user := session.Get("user"); user != nil {
		c.Redirect(http.StatusMovedPermanently, "/mypage")
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func loginPost (c *gin.Context) {
	session := sessions.Default(c)

	username := c.PostForm("username")
	password := c.PostForm("password")

	// ユーザ名からパスワードハッシュを取得
	pass_hash, err := crud.GetPasswordHash(username)
	if err != nil {
		slog.Error(err.Error())
		c.HTML(http.StatusAccepted, "login.html", gin.H{"error": "ユーザ名が不正です."})
		return
	}

	// パスワードを比較
	err = crypto.CompareHashAndPassword(pass_hash, password)
	if err != nil {
		slog.Error(err.Error())
		c.HTML(http.StatusAccepted, "login.html", gin.H{"error": "パスワードが不正です."})
		return
	}

	// ユーザIDを取得
	userid, err := crud.GetUserId(username)
	if err != nil {
		slog.Error(err.Error())
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": "内部サーバーエラーが発生しました."})
		return
	}

	user := user.User{
		UserId: userid,
		UserName: username,
		Password: pass_hash,
		Comportement: user.Comportement{Id: "CP-" + userid },
	}

	user_json, err := json.Marshal(user)
	if err != nil {
		slog.Error(err.Error())
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": "内部サーバーエラーが発生しました."})
		return
	}

	session.Set("user", user_json)
	session.Save()

	c.Redirect(http.StatusMovedPermanently, "/mypage")
}