package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"Goddit/db"
	"Goddit/routes"
)

func main() {

	conn, err := db.InitDB()
	if err != nil {
		return
	}
	router := gin.Default()
	routes.InitPostsRoutes(conn, router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server")
	}
}
