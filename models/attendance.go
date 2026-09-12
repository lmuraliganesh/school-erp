package models

import (
	"time"
)

type Attendance struct {
	ID        int       `json:"id"`
	StudentID int       `json:"student_id"`
	Date      string    `json:"date"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type MarkAttendanceEntry struct {
	StudentID int    `json:"student_id"`
	Status    string `json:"status"`
}

type MarkAttendanceRequest struct {
	Date     string                `json:"date"`
	ClassID  int                   `json:"class_id"`
	Students []MarkAttendanceEntry `json:"students"`
}
