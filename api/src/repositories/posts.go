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

// SearchID brings a post by its ID and returns it
func (repository Posts) SearchID(postID uint64) (models.Post, error) {
	row, err := repository.db.Query(
		`SELECT p.*, u.username FROM posts p
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

// Search brings all posts from users followed by the given user and his own
func (repository Posts) Search(userID uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(
		`SELECT DISTINCT p.*, u.username FROM posts p
		INNER JOIN users u on u.id = p.author_id
		INNER JOIN followers f on p.author_id = f.user_id
		WHERE u.id = ? OR f.follower_id = ?
		ORDER BY p.id desc`,
		userID,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedOn,
			&post.AuthorUsername,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// Update inserts the post changes made by the owner user into database
func (repository Posts) Update(postID uint64, post models.Post) error {
	dbStatement, err := repository.db.Prepare(
		"UPDATE posts SET title = ?, content = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer dbStatement.Close()

	if _, err = dbStatement.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}

	return nil
}

// Delete handles the deletion of an user from the database
func (repository Posts) Delete(postID uint64) error {
	dbStatement, err := repository.db.Prepare(
		"DELETE FROM posts WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer dbStatement.Close()

	if _, err = dbStatement.Exec(postID); err != nil {
		return err
	}

	return nil
}

// SearchByUser brings all the posts by a specific user
func (repository Posts) SearchByUser(userID uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(
		`SELECT p.*, u.username FROM posts p
		JOIN users u on u.id = p.author_id
		WHERE p.author_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post

		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedOn,
			&post.AuthorUsername,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
