package auth

import (
	"errors"
	"log/slog"
	"net/http"
	"oblivion/pkg/crud"
	"oblivion/pkg/error_hanndler"
	"oblivion/pkg/utils/crypto"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func LoginGet() gin.HandlerFunc {
	return func (c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	}
}

func LoginPost() gin.HandlerFunc {
	return func (c *gin.Context) {
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

		session.Set("userid", userid)
		session.Save()

		c.Redirect(http.StatusMovedPermanently, "/mypage")
	}
}

func RegisterGet() gin.HandlerFunc {
	return func (c *gin.Context)  {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	}
}

func RegisterPost() gin.HandlerFunc {
	return func (c *gin.Context)  {
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
		if errors.As(err, &error_hanndler.AlreadyExsistUserError{}) {
			// リダイレクト
			c.JSON(http.StatusBadRequest, gin.H{"message": "Already registered"})
			return
		} else if err != nil {
			slog.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Insert Error",
				"detail": err.Error(),
			})
			return
		}

		// ユーザIDを取得
		userid, err := crud.GetUserId(username)

		// クッキーにログイン情報をセット
		session.Set("userid", userid)
		session.Save()

		// リダイレクト
		c.Redirect(http.StatusMovedPermanently, "/mypage")
	}
}