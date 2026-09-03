package models

import "time"

type StudentProfile struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	ClassID    int       `json:"class_id"`
	RollNumber string    `json:"roll_number"`
	CreatedAt  time.Time `json:"created_at"`
}

type EnrollStudentRequest struct {
	UserID     int    `json:"user_id"`
	ClassID    int    `json:"class_id"`
	RollNumber string `json:"roll_number"`
}
