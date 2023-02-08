package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/iZarrios/gorm-psql-jwt-demo/models"
	"github.com/iZarrios/gorm-psql-jwt-demo/storage"
	"golang.org/x/crypto/bcrypt"
)

const (
	SECRET_JWT_KEY = "this-is-not-a-secret"
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
	// var user *storage.User
	var user *models.User

	c.BindJSON(&user)

	passwordByte := []byte(user.Password)
	hashedPw, err := bcrypt.GenerateFromPassword(passwordByte, 10)
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
	//no need to conver it into an int (Delte works for strings also)
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
	err = storage.GlobalStore.DeleteUser(uint(idInt))
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
func CookieAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get cookie
		cookie, err := c.Cookie("jwt")
		if err == nil {
			if cookie == "ok" {
				c.Next()
				return
			}
		}
		// Cookie verification failed
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden with no cookie"})
		c.Abort()
	}
}

func LoginUser(c *gin.Context) {
	{
		// var loginUser *storage.User
		var loginUser *models.User

		c.BindJSON(&loginUser)
		var userFromDB *models.User

		err := storage.GlobalStore.DB.First(&userFromDB, "email = ?", loginUser.Email).Error
		if err != nil {

			c.JSON(http.StatusUnprocessableEntity,
				gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
			return
		}
		passwordByte := []byte(loginUser.Password)

		// CompareHashAndPassword compares a bcrypt hashed password with its possible
		// plaintext equivalent. Returns nil on success, or an error on failure.
		err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), passwordByte)

		if err != nil {

			c.JSON(http.StatusUnprocessableEntity,
				gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
			return
		}

		claims := &jwt.RegisteredClaims{
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 600)),
			Issuer:    strconv.Itoa(int(userFromDB.Id)),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		ss, err := token.SignedString([]byte(SECRET_JWT_KEY))
        _ = ss

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"data":    gin.H{"user": userFromDB.UserName},
				"message": "could not login try again",
			})
		}
		//https://github.com/gin-gonic/examples/blob/001f7ac527ee46d6404db92955c69b60311086d8/cookie/main.go
		c.SetCookie("jwt", "ok", 600, "/", "localhost", false, true)

		c.JSON(http.StatusOK,
			gin.H{
				"status":  "success",
				"message": "Logged in successfuly!",
			})

	}
}

func LogoutUser(c *gin.Context) {
	c.SetCookie("jwt", "ok", -1, "/", "localhost", false, true)
    c.JSON(200,"hi")

}
