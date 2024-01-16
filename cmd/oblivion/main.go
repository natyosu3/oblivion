package main

import (
	"github.com/gin-gonic/gin"
	"oblivion/pkg/engine"
)

func main() {
	r := gin.Default()
	r = engine.Engine(r)

	r.Run()
}