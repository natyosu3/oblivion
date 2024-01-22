package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"oblivion/pkg/engine/auth"
	"oblivion/pkg/engine/top"
	"oblivion/pkg/engine/mypage"
)

func Engine(r *gin.Engine) *gin.Engine  {
	r.LoadHTMLGlob("web/templates/*/*.html")

	r.Static("/static", "web/static")

	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	r.Use(sessions.Sessions("session", store))
	

	topGroup := r.Group("/")
	{
		topGroup.GET("/", top.Index())
	}

	authGroup := r.Group("/auth")
	{
		authGroup.GET("/login", auth.LoginGet())
		authGroup.POST("/login", auth.LoginPost())

		authGroup.GET("/register", auth.RegisterGet())
		authGroup.POST("/register", auth.RegisterPost())
	}

	mypageGroup := r.Group("/mypage")
	{
		mypageGroup.GET("/", mypage.MypageTop())
	}

	return r
}