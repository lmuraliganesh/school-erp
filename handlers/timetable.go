package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"school-erp/database"
	"school-erp/models"
)

func CreateTimetableHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CreateTimetableRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var newID int
	query := `INSERT INTO timetables (class_id, subject_id, day_of_week, start_time, end_time) VALUES ($1,$2,$3,$4,$5) RETURNING id`
	err = database.DB.QueryRow(context.Background(), query, req.ClassID, req.SubjectID, req.DayOfWeek, req.StartTime, req.EndTime).Scan(&newID)
	if err != nil {
		http.Error(w, "Failed to create timetable entry (check class_id/subject_id are valid)", http.StatusBadRequest)
		return
	}

	response := models.Timetable{
		ID:        newID,
		ClassID:   req.ClassID,
		SubjectID: req.SubjectID,
		DayOfWeek: req.DayOfWeek,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
func GetTimetableHandler(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, class_id, subject_id, day_of_week, start_time, end_time FROM timetables ORDER BY id`
	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		http.Error(w, "Failed to fetch timetable", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var timetable []models.Timetable
	for rows.Next() {
		var t models.Timetable
		err := rows.Scan(&t.ID, &t.ClassID, &t.SubjectID, &t.DayOfWeek, &t.StartTime, &t.EndTime)
		if err != nil {
			http.Error(w, "Failed to read timetable data", http.StatusInternalServerError)
			return
		}
		timetable = append(timetable, t)
	}

	// 3. Respond with the list:
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timetable)
}
