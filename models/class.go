package models

import "time"

type Class struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Section   string    `json:"section"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateClassRequest struct {
	Name    string `json:"name"`
	Section string `json:"section"`
}

type Subject struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateSubjectRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
