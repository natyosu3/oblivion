package component

import (
	"encoding/json"
	"net/http"
	"oblivion/pkg/crud"
	"oblivion/pkg/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AddComponentGet() gin.HandlerFunc {
	return func (c *gin.Context)  {
		c.HTML(200, "add.html", gin.H{
			"IsAuthenticated": true,
		})
	}
}

func AddComponentPost() gin.HandlerFunc {
	return func (c *gin.Context)  {
		session := sessions.Default(c)

		if session.Get("user") == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
			return
		}

		user := user.User{}
		if err := json.Unmarshal(session.Get("user").([]byte), &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		elementName := c.PostForm("name")
		content := c.PostForm("content")

		err := crud.InsertElement(user.UserId, elementName, content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.HTML(200, "add.html", gin.H{
			"message": "success",
			"IsAuthenticated": true,
		})
	}
}

func DeleteComponentPost() gin.HandlerFunc {
	return func (c *gin.Context)  {
		session := sessions.Default(c)

		if session.Get("user") == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
			return
		}

		user := user.User{}
		if err := json.Unmarshal(session.Get("user").([]byte), &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		elementId := c.PostForm("id")

		err := crud.DeleteElement(user.UserId, elementId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.Redirect(http.StatusMovedPermanently, "/component/list")
	}
}

func ListComponentGet() gin.HandlerFunc {
	return func (c *gin.Context)  {
		session := sessions.Default(c)

		if session.Get("user") == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/auth/login")
			return
		}

		user := user.User{}
		if err := json.Unmarshal(session.Get("user").([]byte), &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		elements, err := crud.ListElement(user.UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.HTML(200, "list.html", gin.H{
			"elements": elements,
			"IsAuthenticated": true,
		})
	}
}