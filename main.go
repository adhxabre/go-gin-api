package main

import (
	"go-gin-api/pkg/database"
	"go-gin-api/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("[ERR] .env failed to load, is it already configured/available?")
	}

	PORT := os.Getenv("PORT")

	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	database.DBConn()
	database.DBMigration()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	routes.RouteInit(r)

	servErr := r.Run("localhost:" + PORT)
	if servErr != nil {
		panic("[ERR] failed to start application due to: " + servErr.Error())
	}
}
