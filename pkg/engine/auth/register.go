package auth

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"encoding/json"
	"log/slog"
	"net/http"
	"oblivion/pkg/crud"
	"oblivion/pkg/error_handler"
	"oblivion/pkg/user"
	"oblivion/pkg/utils/crypto"
)

func registerGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func registerPost(c *gin.Context) {
	session := sessions.Default(c)

	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	// パスワードをハッシュ化
	hash, err := crypto.PasswordEncrypt(password)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Encrypt Error"})
		return
	}

	// ユーザー登録
	err = crud.InsertUser(username, email, hash)
	if errors.As(err, &error_handler.AlreadyExsistUserError{}) {
		// リダイレクト
		c.JSON(http.StatusBadRequest, gin.H{"message": "Already registered"})
		return
	} else if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Insert Error",
			"detail":  err.Error(),
		})
		return
	}

	// ユーザIDを取得
	userid, err := crud.GetUserId(username)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Get User ID Error"})
		return
	}

	user := user.User{
		UserId:       userid,
		UserName:     username,
		EmailAddress: email,
		Comportement: user.Comportement{Id: "CP-" + userid},
	}

	user_json, err := json.Marshal(user)
	if err != nil {
		slog.Error(err.Error())
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": "内部サーバーエラーが発生しました."})
		return
	}

	// クッキーにログイン情報をセット
	session.Set("user", user_json)
	session.Save()

	// リダイレクト
	c.Redirect(http.StatusSeeOther, "/mypage")

}
