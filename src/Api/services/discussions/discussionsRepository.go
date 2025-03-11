package discussions

import (
	"Goddit/models"

	"github.com/rqlite/gorqlite"
)

func GetRecentDiscussions(db *gorqlite.Connection) (*[]models.Discussion, error) {
	return &[]models.Discussion{}, nil
}

func GetRelatedDiscussions(kwds []string, db *gorqlite.Connection) (*[]models.Discussion, error) {
	return &[]models.Discussion{}, nil
}

func SearchDiscussions(kwds []string, db *gorqlite.Connection) (*[]models.Discussion, error) {
	return &[]models.Discussion{}, nil
}

func GetOwnDiscussions(user *models.User, db *gorqlite.Connection) (*[]models.Discussion, error) {
	return &[]models.Discussion{}, nil
}
