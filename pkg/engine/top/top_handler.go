package top

import (
	"github.com/gin-gonic/gin"
)


func Index() gin.HandlerFunc {
	return func (c *gin.Context)  {
		c.HTML(200, "index.html", gin.H{
			"IsAuthenticated": true,
		})
	}
}