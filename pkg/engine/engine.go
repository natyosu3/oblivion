package engine

import (
	"log/slog"
	"oblivion/pkg/engine/auth"
	"oblivion/pkg/engine/component"
	"oblivion/pkg/engine/mypage"
	"oblivion/pkg/engine/top"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	REDIS_HOST     string
	REDIS_PASSWORD string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("Error loading .env file")
	}

	REDIS_HOST = os.Getenv("REDIS_HOST")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
}

func Engine(r *gin.Engine) *gin.Engine {
	r.LoadHTMLGlob("web/templates/*/*.html")
	r.Static("/static", "web/static")

	store, err := redis.NewStore(10, "tcp", REDIS_HOST, REDIS_PASSWORD, []byte("secret"))
	if err != nil {
		slog.Error(err.Error())
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
		authGroup.GET("/logout", auth.LogoutGet())

		authGroup.GET("/discord/login", auth.DiscordLoginGet())
		authGroup.GET("/callback", auth.DiscordCallbackGet())

		authGroup.GET("/register", auth.RegisterGet())
		authGroup.POST("/register", auth.RegisterPost())
	}

	componentGroup := r.Group("/component")
	{
		componentGroup.GET("/add", component.AddComponentGet())
		componentGroup.POST("/add", component.AddComponentPost())
		componentGroup.POST("/delete", component.DeleteComponentPost())
		componentGroup.GET("/list", component.ListComponentGet())
		componentGroup.POST("/check", component.CheackComponentPost())
		componentGroup.GET("/edit/:id", component.EditComponentGet())
		componentGroup.POST("/edit/:id", component.EditComponentPost())
	}

	mypageGroup := r.Group("/mypage")
	{
		mypageGroup.GET("/", mypage.MypageTop())
	}

	return r
}
