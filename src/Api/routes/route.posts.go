package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rqlite/gorqlite"
)

func InitPostsRoutes(db *gorqlite.Connection, router *gin.Engine) {

	// This is initial routing. We should create the high level architecture of routing and divide in logical blocks (groups)

	r := router.Group("/v1/posts")
	r.GET("", func(ctx *gin.Context) {
		results, err := db.QueryOne("SELECT * FROM posts")
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

		ctx.JSON(http.StatusOK, posts)
	})
}
