package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"school-erp/database"
	"school-erp/models"
)

func EnrollStudentHandler(w http.ResponseWriter, r *http.Request) {
	var req models.EnrollStudentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var newID int
	query := `INSERT INTO student_profiles(user_id, class_id, roll_number) VALUES ($1,$2,$3) RETURNING id`
	err = database.DB.QueryRow(context.Background(), query, req.UserID, req.ClassID, req.RollNumber).Scan(&newID)
	if err != nil {
		http.Error(w, "Failed to enroll student (check user_id/class_id are valid)", http.StatusBadRequest)
		return
	}

	response := models.StudentProfile{

		ID:         newID,
		UserID:     req.UserID,
		ClassID:    req.ClassID,
		RollNumber: req.RollNumber,
	}
	w.Header().Set("Content type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
