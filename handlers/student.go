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

func GetStudentsHandler(w http.ResponseWriter, r *http.Request) {
	query := `
	SELECT 
		student_profiles.id,
		users.name,
		users.email,
		classes.name AS class_name,
		classes.section,
		student_profiles.roll_number
	FROM student_profiles
	JOIN users ON student_profiles.user_id = users.id
	JOIN classes ON student_profiles.class_id = classes.id
	ORDER BY student_profiles.id
	`

	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		http.Error(w, "Failed to fetch students", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []models.StudentDirectoryEntry
	for rows.Next() {
		var s models.StudentDirectoryEntry
		err := rows.Scan(&s.ID, &s.Name, &s.Email, &s.ClassName, &s.Section, &s.RollNumber)
		if err != nil {
			http.Error(w, "Failed to parse student data", http.StatusInternalServerError)
			return
		}
		students = append(students, s)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Error reading student records", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(students); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
