package models

import "time"

type Announcement struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateAnnouncementRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
