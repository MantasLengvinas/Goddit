package data

import (
	"context"
	"log"

	"github.com/rqlite/gorqlite"
)

type PostRepository struct {
	db *gorqlite.Connection
}

func InitPostRepository(db *gorqlite.Connection) *PostRepository {
	return &PostRepository{db: db}
}

func (repo *PostRepository) GetAllPosts(ctx context.Context) []map[string]any {
	results, err := repo.db.QueryOne("SELECT * FROM posts")
	if err != nil {
		log.Printf("Query error")
		return nil
	}

	var posts []map[string]any
	for results.Next() {
		row, err := results.Map()
		if err != nil {
			log.Printf("Row reading error")
			return nil
		}
		posts = append(posts, row)
	}
	return posts
}

func (repo *PostRepository) GetPost(ctx context.Context, id string) map[string]any {
	// queury parametrized to avoid sql injection
	results, err := repo.db.QueryParameterized([]gorqlite.ParameterizedStatement{
		{
			Query: "SELECT * FROM posts WHERE id = ? LIMIT 1",
			Arguments: []any { id },
		},
	})

	if err != nil {
		log.Printf("Query error")
		return nil
	}

	// we only need one post, so take a slice of the results (which should be one post anyhow)
	post := results[0]

	if post.Next() {
		row, err := post.Map()
		if err != nil {
			log.Printf("Row reading error")
			return nil
		}
		return row
	}
	return nil
}
