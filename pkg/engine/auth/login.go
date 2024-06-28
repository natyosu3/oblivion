package auth

import (
	"log/slog"
	"net/http"
	"oblivion/pkg/crud"
	"oblivion/pkg/discord"
	"oblivion/pkg/session"
	"oblivion/pkg/user"
	"oblivion/pkg/utils/crypto"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func loginGet(c *gin.Context) {
	data := session.Default(c, "session", &user.User{}).Get(c)

	if data != nil {
		c.Redirect(http.StatusSeeOther, "/mypage")
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func loginPost(c *gin.Context) {
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

	se_data := user.User{
		UserId:       userid,
		UserName:     username,
		Comportement: user.Comportement{Id: "CP-" + userid},
	}

	// セッションを設定(cookieにセット)
	session.Default(c, "session", &user.User{}).Set(c, se_data)

	c.Redirect(http.StatusMovedPermanently, "/mypage")
}

func discordLoginGet(c *gin.Context) {
	c.Redirect(http.StatusSeeOther, discord.Oauth2_URL)
}

func discordCallbackGet(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(400, gin.H{"message": "code is empty"})
		return
	}
	tmp := "client_id=" + discord.Client_Id + "&client_secret=" + discord.Client_Secret + "&grant_type=authorization_code&code=" + code + "&redirect_uri=https://oblivion.natyosu.com/auth/callback"

	payload := []byte(tmp)
	token, err := discord.Oauth2(payload)
	if err != nil {
		c.JSON(500, gin.H{"message": "Error oauth2"})
		return
	}
	// セッションに resValue を保存
	session := sessions.Default(c)
	session.Set("token", token)
	session.Save()

	c.Redirect(302, "/mypage")
}
