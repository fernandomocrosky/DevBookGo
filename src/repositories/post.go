package repositories

import (
	"database/sql"

	"github.com/fernandomocrosky/DevBookGo/src/models"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (repo *PostRepository) CreatePost(post models.Posts) (uint64, error) {
	statement, err := repo.db.Prepare(
		`INSERT INTO posts (title, content, user_id) VALUES (?, ?, ?)`,
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.UserId)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}
