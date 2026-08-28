package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"school-erp/database"
	"school-erp/models"
	"school-erp/utils"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Query all users:
	query := `SELECT id, name, email, role FROM users ORDER BY id`
	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// 2. Loop through rows, build a slice:
	var users []models.UserResponse
	for rows.Next() {
		var u models.UserResponse
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Role)
		if err != nil {
			http.Error(w, "Failed to read user data", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	// 3. Respond with the list:
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 2. Validate Role is one of "admin", "teacher", "student"
	if req.Role != "admin" && req.Role != "teacher" && req.Role != "student" {
		http.Error(w, "Invalid role", http.StatusBadRequest)
		return
	}

	// 3. Hash req.Password using utils.HashPassword
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to process password", http.StatusInternalServerError)
		return
	}

	// 4. Insert into DB, get new ID back:
	var newID int
	query := `INSERT INTO users (name, email, password_hash, role) VALUES ($1,$2,$3,$4) RETURNING id`
	err = database.DB.QueryRow(context.Background(), query, req.Name, req.Email, hashedPassword, req.Role).Scan(&newID)
	if err != nil {
		http.Error(w, "Failed to create user (email may already exist)", http.StatusConflict)
		return
	}

	// 5. Build models.UserResponse{ID: newID, Name: req.Name, Email: req.Email, Role: req.Role}
	response := models.UserResponse{
		ID:    newID,
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}

	// 6. w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
