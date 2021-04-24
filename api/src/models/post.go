package models

import "time"

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
