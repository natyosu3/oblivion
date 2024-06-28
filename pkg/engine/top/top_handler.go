package top

import (
	"oblivion/pkg/model"
	"oblivion/pkg/session"

	"github.com/gin-gonic/gin"
)

func Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := session.Default(c, "session", &model.Session_model{}).Get(c)
		if data == nil {
			c.HTML(200, "index.html", gin.H{
				"IsAuthenticated": false,
			})
			return
		}
		c.HTML(200, "index.html", gin.H{
			"IsAuthenticated": true,
		})
	}
}
