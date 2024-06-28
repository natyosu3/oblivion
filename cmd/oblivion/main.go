package main

import (
	"log/slog"
	"oblivion/pkg/engine"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Info("【GIN】This is a Production Environment")
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()
	r = engine.Engine(r)

	r.Run()
}
