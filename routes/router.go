package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iZarrios/gorm-psql-jwt-demo/controllers"
)

func SetupRouter(rg *gin.RouterGroup) {
    r := rg.Group("/")

	r.GET("/ping", controllers.Ping)
}
