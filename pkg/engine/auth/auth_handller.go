package auth

import (
	"errors"
	"oblivion/pkg/error_hanndler"
	"log/slog"
	"net/http"
	"oblivion/pkg/utils/crypto"
	"oblivion/pkg/crud"
	"github.com/gin-gonic/gin"
)


func LoginGet() gin.HandlerFunc {
	return func (c *gin.Context)  {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	}
}

func LoginPost() gin.HandlerFunc {
	return func (c *gin.Context)  {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	}
}

func RegisterGet() gin.HandlerFunc {
	return func (c *gin.Context)  {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	}
}

func RegisterPost() gin.HandlerFunc {
	return func (c *gin.Context)  {
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
		if errors.As(err, &error_hanndler.AlreadyExsistUserError{}) {
			// クッキーにエラーをセット

			// リダイレクト
			c.Redirect(http.StatusTemporaryRedirect, "/register")
			return
		} else if err != nil {
			slog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"message": "Insert Error"})
			return
		}

		slog.Info("User created")

		// クッキーにログイン情報をセット

		// リダイレクト
		c.Redirect(http.StatusTemporaryRedirect, "/mypage")
	}
}