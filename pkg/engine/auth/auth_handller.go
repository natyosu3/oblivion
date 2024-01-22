package auth

import (
	"github.com/gin-gonic/gin"
)

func LoginGet() gin.HandlerFunc {
	return loginGet
}

func LoginPost() gin.HandlerFunc {
	return loginPost
}

func RegisterGet() gin.HandlerFunc {
	return registerGet()
}

func RegisterPost() gin.HandlerFunc {
	return registerPost()
}

func LogoutGet() gin.HandlerFunc {
	return logoutGet()
}