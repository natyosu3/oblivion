package component

import (
	"github.com/gin-gonic/gin"
)

func AddComponentGet() gin.HandlerFunc {
	return addComponentGet
}

func AddComponentPost() gin.HandlerFunc {
	return addComponentPost
}

func DeleteComponentPost() gin.HandlerFunc {
	return deleteComponentPost
}

func ListComponentGet() gin.HandlerFunc {
	return listComponentGet
}

func CheackComponentPost() gin.HandlerFunc {
	return cheackComponentPost
}

func EditComponentGet() gin.HandlerFunc {
	return editComponentGet
}

func EditComponentPost() gin.HandlerFunc {
	return editComponentPost
}