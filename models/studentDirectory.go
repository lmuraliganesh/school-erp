package models

type StudentDirectoryEntry struct {
	UserID     int    `json:"user_id"`
	ClassID    int    `json:"class_id"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	ClassName  string `json:"class_name"`
	Section    string `json:"section"`
	RollNumber string `json:"roll_number"`
}
