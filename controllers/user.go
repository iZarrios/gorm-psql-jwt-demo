package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iZarrios/gorm-psql-jwt-demo/storage"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {

	users, err := storage.GlobalStore.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    "success",
		"message": users,
	})

}
func CreateUser(c *gin.Context) {
	var user *storage.User

	c.BindJSON(&user)

	password := []byte(user.Password)
	hashedPw, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {

		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  "failed",
				"message": err,
			})
		return
	}
	user.Password = string(hashedPw)

	err = storage.GlobalStore.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"status":  "failed",
				"message": err,
			})
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"status":  "success",
			"message": "Created user successfuly!",
			"data":    user,
		})

}

func DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity,
			gin.H{
				"status":  "failed",
				"message": "error with input",
				"data":    nil,
			})
		return
	}
	err = storage.GlobalStore.DeleteUser(idInt)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity,
			gin.H{
				"status":  "failed",
				"message": "ID not found",
				"data":    nil,
			})
		return
	}
	c.JSON(http.StatusOK,
		gin.H{
			"status":  "success",
			"message": "Successfuly deleted user",
			"data":    idInt,
		})
}
