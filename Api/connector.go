package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rqlite/gorqlite"
)

func main() {

	router := gin.Default()

	router.GET("/posts", func(context *gin.Context) {
		conn, err := gorqlite.Open("http://")
		if err != nil {
			log.Printf("DB connection error")
			return
		}

		results, err := conn.QueryOne("SELECT * FROM posts")
		if err != nil {
			log.Printf("Query error")
			return
		}

		var posts []map[string]interface{}
		for results.Next() {
			row, err := results.Map()
			if err != nil {
				log.Printf("Row reading error")
				return
			}
			posts = append(posts, row)
		}

		context.JSON(http.StatusOK, posts)

	})

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server")
	}
}