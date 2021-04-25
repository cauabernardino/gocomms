package repositories

import (
	"api/src/models"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

// NewPostsRepository handles link between the post struct and database
func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

// Create inserts a post to the database
func (repository Posts) Create(post models.Post) (uint64, error) {
	dbStatement, err := repository.db.Prepare(
		"INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer dbStatement.Close()

	dbExecution, err := dbStatement.Exec(
		post.Title,
		post.Content,
		post.AuthorID,
	)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := dbExecution.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}

// SearchID searches by a post by its ID
func (repository Posts) SearchID(postID uint64) (models.Post, error) {
	row, err := repository.db.Query(`
		SELECT p.*, u.username FROM posts p
		INNER JOIN users u ON u.id = p.author_id 
		WHERE p.id = ?`,
		postID,
	)
	if err != nil {
		return models.Post{}, err
	}
	defer row.Close()

	var post models.Post
	if row.Next() {
		if err = row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedOn,
			&post.AuthorUsername,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}
