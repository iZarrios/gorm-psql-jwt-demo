package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {


	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "pong",
	})

}


