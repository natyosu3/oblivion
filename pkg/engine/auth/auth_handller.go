package auth

import (
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
		c.HTML(200, "register.html", gin.H{})
	}
}