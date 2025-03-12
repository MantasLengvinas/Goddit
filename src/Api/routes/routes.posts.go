package routes

import (
	posts "Goddit/services/posts"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rqlite/gorqlite"
)

func InitPostsRoutes(db *gorqlite.Connection, router *gin.Engine) {

	// This is initial routing. We should create the high level architecture of routing and divide in logical blocks (groups)
	repo := posts.InitPostRepository(db)
	r := router.Group("/v1/posts")

	// GET /v1/posts -> return all, index
	r.GET("", func(ctx *gin.Context) {
		posts := repo.GetAllPosts(ctx)

		ctx.JSON(http.StatusOK, posts)
	})

	// GET /v1/posts/:id -> get post by id
	r.GET("/:id", func(ctx *gin.Context) {
		// currently id is understood as string, not sure if should be cast down to int later for cohesion
		id := ctx.Param("id")
		post := repo.GetPost(ctx, id)

		ctx.JSON(http.StatusOK, post)
	})
}
