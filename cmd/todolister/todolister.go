package main

import (
	"github.com/calz10/todolister/db"
	"github.com/calz10/todolister/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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
