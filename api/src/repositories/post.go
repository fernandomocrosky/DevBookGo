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

func (repo *PostRepository) GetPosts(userId uint64) ([]models.Posts, error) {
	rows, err := repo.db.Query(`
		SELECT DISTINCT 
			p.*, u.nick
		FROM posts p
		INNER JOIN users u ON u.id = p.user_id 
		INNER JOIN followers f ON p.user_id = f.user_id
		WHERE u.id = ? or f.follower_id = ?
		ORDER BY  1 DESC
	`, userId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Posts

	for rows.Next() {
		var post models.Posts
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserId, &post.Likes, &post.CreatedAt, &post.UserNick); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (repo *PostRepository) GetPostById(postId uint64) (models.Posts, error) {
	row, err := repo.db.Query(`
		SELECT 
			p.*, u.nick
		FROM posts p 
		INNER JOIN users u ON u.id = p.user_id WHERE p.id = ?
	`, postId)
	if err != nil {
		return models.Posts{}, err
	}
	defer row.Close()

	var post models.Posts
	if row.Next() {
		if err := row.Scan(&post.ID, &post.Title, &post.Content, &post.UserId, &post.Likes, &post.CreatedAt, &post.UserNick); err != nil {
			return models.Posts{}, err
		}
	}

	return post, nil
}

func (repo *PostRepository) UpdatePost(postId uint64, post models.Posts) error {
	statement, err := repo.db.Prepare(`
		UPDATE
			posts
		SET
			title = ?, content = ? 
		WHERE
			id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(post.Title, post.Content, postId); err != nil {
		return err
	}

	return nil
}

func (repo *PostRepository) DeletePost(postId uint64) error {
	statement, err := repo.db.Prepare(`
		DELETE FROM posts WHERE id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(postId); err != nil {
		return err
	}
	return nil
}

func (repo *PostRepository) GetPostsByUserId(userId uint64) ([]models.Posts, error) {
	rows, err := repo.db.Query(`
		SELECT p.*, u.nick
		FROM posts p
		INNER JOIN users u ON u.id = p.user_id WHERE p.user_id = ?
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Posts = []models.Posts{}

	for rows.Next() {
		var post models.Posts
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.UserId,
			&post.Likes,
			&post.CreatedAt,
			&post.UserNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repo *PostRepository) LikePost(postId uint64) error {
	statement, err := repo.db.Prepare(
		`
			UPDATE 
				posts 
			SET
				likes = likes + 1
			WHERE id = ?
		`,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(postId); err != nil {
		return err
	}

	return nil
}

func (repo *PostRepository) UnlikePost(postId uint64) error {
	statement, err := repo.db.Prepare(`
		UPDATE posts
		SET 
			likes = 
				CASE 
					WHEN likes > 0 THEN likes - 1 
					ELSE likes 
				END
		WHERE id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(postId); err != nil {
		return err
	}

	return nil
}
