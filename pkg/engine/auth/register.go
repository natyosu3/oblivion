package auth

import (
	"errors"

	"github.com/gin-gonic/gin"

	"log/slog"
	"net/http"
	"oblivion/pkg/crud"
	"oblivion/pkg/error_handler"
	"oblivion/pkg/model"
	"oblivion/pkg/session"
	"oblivion/pkg/utils/crypto"
)

func registerGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func registerPost(c *gin.Context) {
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

	// ユーザ情報をセッションに保存
	se_data := model.Session_model{
		SessionId: "",
		CookieKey: "session",
		User: model.User{
			UserId:       userid,
			UserName:     username,
			Comportement: model.Comportement{Id: "CP-" + userid},
		},
		Token: "",
	}

	session.Default(c, "session", &model.Session_model{}).Set(c, se_data)

	// リダイレクト
	c.Redirect(http.StatusSeeOther, "/mypage")

}
