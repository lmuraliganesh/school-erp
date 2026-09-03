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
