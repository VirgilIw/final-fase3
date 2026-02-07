package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/virgilIw/final-fase3/docs"
	"github.com/virgilIw/final-fase3/internal/config"
	"github.com/virgilIw/final-fase3/internal/router"
)

// @title						social media Backend
// @version						1.0
// @description					Social media Backend RESTful API
// @host						localhost:8080
// @BasePath					/
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description					Type "Bearer" followed by a space and JWT token.
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to Load env")
		return
	}
	db, err := config.InitDB()
	if err != nil {
		log.Println("Failed to Connect to Database")
		return
	}
	rdb := config.InitRds()
	defer rdb.Close()

	app := gin.Default()

	router.Init(app, db, rdb)

	app.Run()
}
