package auth

import (
	"fmt"
	"oblivion/pkg/utils/crypto"
	"oblivion/pkg/crud"
	"github.com/gin-gonic/gin"
)


func Login() gin.HandlerFunc {
	return func (c *gin.Context)  {
		c.HTML(200, "login.html", gin.H{})
	}
}

func RegisterGet() gin.HandlerFunc {
	return func (c *gin.Context)  {
		c.HTML(200, "register.html", gin.H{})
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
			fmt.Println(err)
			c.JSON(500, gin.H{"message": "Internal Server Error"})
		}

		// ユーザー登録
		err = crud.InsertUser(username, email, hash)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"message": "Internal Server Error"})
		}

		c.JSON(200, gin.H{"message": "OK"})
	}
}