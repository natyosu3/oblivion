package main

import (
	"github.com/gin-gonic/gin"
	"oblivion/pkg/engine"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r = engine.Engine(r)

	r.Run()
}