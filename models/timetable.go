package models

import "time"

type Timetable struct {
	ID        int       `json:"id"`
	ClassID   int       `json:"class_id"`
	SubjectID int       `json:"subject_id"`
	DayOfWeek string    `json:"day_of_week"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateTimetableRequest struct {
	ClassID   int    `json:"class_id"`
	SubjectID int    `json:"subject_id"`
	DayOfWeek string `json:"day_of_week"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
