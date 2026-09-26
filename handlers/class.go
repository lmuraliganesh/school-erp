package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"school-erp/database"
	"school-erp/models"
)

func CreateClassHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CreateClassRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 1. Insert into DB, get new ID back:
	var newID int
	query := `INSERT INTO classes (name, section) VALUES ($1,$2) RETURNING id`
	err = database.DB.QueryRow(context.Background(), query, req.Name, req.Section).Scan(&newID)
	if err != nil {
		http.Error(w, "Failed to create class", http.StatusInternalServerError)
		return
	}

	// 2. Build response
	response := models.Class{
		ID:      newID,
		Name:    req.Name,
		Section: req.Section,
	}

	// 3. Respond
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
func GetClassesHandler(w http.ResponseWriter, r *http.Request) {
	// SELECT id, name, section, created_at FROM classes ORDER BY id
	// Query, loop, Scan into []models.Class, respond
	query := `SELECT id, name, section, created_at FROM classes ORDER BY id`
	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		http.Error(w, "failed to fetch classes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var classesList []models.Class

	for rows.Next() {
		var c models.Class
		err := rows.Scan(&c.ID, &c.Name, &c.Section, &c.CreatedAt)
		if err != nil {
			http.Error(w, "failed to parse the class record", http.StatusInternalServerError)
			return
		}
		classesList = append(classesList, c)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(classesList)
}

func CreateSubjectHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CreateSubjectRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Insert into DB, handling potential unique constraint violations on code
	var newID int
	query := `INSERT INTO subjects (name, code) VALUES ($1,$2) RETURNING id`
	err = database.DB.QueryRow(context.Background(), query, req.Name, req.Code).Scan(&newID)
	if err != nil {
		http.Error(w, "Failed to create subject (code may already exist)", http.StatusConflict)
		return
	}

	response := models.Subject{
		ID:   newID,
		Name: req.Name,
		Code: req.Code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func GetSubjectsHandler(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, name, code, created_at FROM subjects ORDER BY id`
	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		http.Error(w, "Failed to fetch the subject", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var subjectlist []models.Subject
	for rows.Next() {

		var s models.Subject
		err := rows.Scan(&s.ID, &s.Name, &s.Code, &s.CreatedAt)
		if err != nil {
			http.Error(w, "failed to parse the subject record", http.StatusInternalServerError)
			return
		}
		subjectlist = append(subjectlist, s)

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(subjectlist)

}
