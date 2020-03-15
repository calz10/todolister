package main

import (
	"fmt"
	"log"
	"os"

	"github.com/calz10/todolister/db"
	"github.com/calz10/todolister/env/constants"

	"github.com/calz10/todolister/routes"
	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	envMode := os.Getenv(constants.ENV_MODE)
	envKey := os.Getenv(constants.JWT_KEY)
	envFile := ".env." + envMode
	fmt.Print(envFile, "<<<==== env file")
	fmt.Print(envKey, "<<<==== env ke")
	if err := godotenv.Load(envFile); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	db := db.Start()
	router := gin.Default()

	dbService := routes.DbService{DB: db}
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Access-Control-Allow-Origin", "Authorization", "Origin", "Content-Type"}

	dbService.InitializeApi(router)
	router.Use(cors.New(config))

	router.Run()
}
