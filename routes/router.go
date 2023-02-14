package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iZarrios/gorm-psql-jwt-demo/controllers"
)

func SetupRouter(rg *gin.RouterGroup) {
	r := rg.Group("/")
	auth := rg.Group("/auth")
	admin := rg.Group("/admin")

	r.GET("/ping", controllers.Ping)

	admin.Use(controllers.CookieAuth())
	admin.GET("/users", controllers.GetUsers)
	admin.DELETE("/users/:id", controllers.DeleteUser)

	auth.POST("/register", controllers.CreateUser)
	auth.POST("/login", controllers.LoginUser)
	auth.GET("/logout", controllers.LogoutUser)
}
