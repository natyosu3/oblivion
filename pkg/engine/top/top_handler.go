package top

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)


func Index() gin.HandlerFunc {
	return func (c *gin.Context)  {
		session := sessions.Default(c)
		if user := session.Get("user"); user == nil {
			c.HTML(200, "index.html", gin.H{
				"IsAuthenticated": false,
			})
		}
		c.HTML(200, "index.html", gin.H{
			"IsAuthenticated": true,
		})
	}
}