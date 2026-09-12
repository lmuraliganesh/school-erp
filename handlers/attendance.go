package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"school-erp/database"
	"school-erp/models"

	"github.com/go-chi/chi/v5"
)

func MarkAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	var req models.MarkAttendanceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for _, entry := range req.Students {
		query := `INSERT INTO attendance(student_id, date, status) VALUES ($1,$2,$3)`
		_, err := database.DB.Exec(context.Background(), query, entry.StudentID, req.Date, entry.Status)
		if err != nil {
			http.Error(w, "Failed to mark attendance for student", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Attendance marked successfully",
		"count":   len(req.Students),
	})

}
func GetStudentAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	studentID := chi.URLParam(r, "id")

	query := `SELECT id, student_id, date, status, created_at FROM attendance WHERE student_id = $1 ORDER BY date`
	rows, err := database.DB.Query(context.Background(), query, studentID)
	if err != nil {
		http.Error(w, "Failed to fetch attendance records", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var attendanceList []models.Attendance

	for rows.Next() {
		var a models.Attendance
		var dateVal time.Time
		err := rows.Scan(&a.ID, &a.StudentID, &dateVal, &a.Status, &a.CreatedAt)
		if err != nil {
			http.Error(w, "Failed to parse attendance record", http.StatusInternalServerError)
			return
		}
		a.Date = dateVal.Format("2006-01-02")
		attendanceList = append(attendanceList, a)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Error processing attendance rows", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(attendanceList)
}
