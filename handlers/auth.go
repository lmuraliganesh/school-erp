package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"school-erp/database"
	"school-erp/models"
	"school-erp/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 1. Look up the user by email
	var user models.Users
	query := `SELECT id, name, email, password_hash, role FROM users WHERE email = $1`
	err = database.DB.QueryRow(context.Background(), query, req.Email).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// 2. Check the password
	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// 3. Generate the token
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// 4. Respond with the token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
