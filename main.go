package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iZarrios/gorm-psql-jwt-demo/models"
	"github.com/iZarrios/gorm-psql-jwt-demo/routes"
	"github.com/iZarrios/gorm-psql-jwt-demo/storage"
)

func main() {

	myDbConfig := storage.DBConfig{
		Host:     "localhost",
		Port:     "5432",
		Password: "demo123",
		User:     "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	}

	// singal instance that we are connected to all the time
	store := &storage.PostgresStore{}
	storage.ConnectDB(&myDbConfig, store)
    storage.GlobalStore = store

	fmt.Println("Connected to db...")
    models.MigrateUser(storage.GlobalStore.DB)
    fmt.Println("Successfully migrated the user table to the db...")

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	v1 := r.Group("/api/v1")

	routes.SetupRouter(v1)
	r.Run("localhost:8000")


}
