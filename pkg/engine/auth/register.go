package auth

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"oblivion/pkg/crud"
	"oblivion/pkg/error_handler"
	"oblivion/pkg/user"
	"oblivion/pkg/utils/crypto"
	"errors"
	"log/slog"
)

func registerGet() gin.HandlerFunc {
	return func (c *gin.Context)  {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	}
}

func registerPost() gin.HandlerFunc {
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
		if errors.As(err, &error_handler.AlreadyExsistUserError{}) {
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

		user := user.User{
			UserId: userid,
			UserName: username,
			EmailAddress: email,
			Password: hash,
			Comportement: user.Comportement{Id: "CP-" + userid },
		}

		// クッキーにログイン情報をセット
		session.Set("user", user)
		session.Save()

		// リダイレクト
		c.Redirect(http.StatusMovedPermanently, "/mypage")
	}
}