package models

import (
	"errors"
	"strings"
	"time"
)

// Post represents the structure of a post
type Post struct {
	ID             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorID       uint64    `json:"author_id,omitempty"`
	AuthorUsername string    `json:"author_username,omitempty"`
	Likes          uint64    `json:"likes"`
	CreatedOn      time.Time `json:"created_on,omitempty"`
}

// Prepare calls the validation and formatting functions for
// the Post struct
func (post *Post) Prepare() error {
	if err := post.validate(); err != nil {
		return err
	}

	post.format()

	return nil
}

// validate the post struct
func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("title field is required")
	}

	if post.Content == "" {
		return errors.New("content field is required")
	}

	return nil
}

// format the post struct
func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
