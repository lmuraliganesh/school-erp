package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"school-erp/database"
	appmw "school-erp/middleware"
	"school-erp/models"
	"school-erp/utils"
)

func CreateAnnouncementHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CreateAnnouncementRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	claims, ok := r.Context().Value(appmw.UserContextKey).(*utils.Claims)
	if !ok {
		http.Error(w, "Unauthourized", http.StatusUnauthorized)
		return
	}
	var newID int
	query := `INSERT INTO announcements(title, content, author_id) VALUES ($1,$2,$3) RETURNING id`
	err = database.DB.QueryRow(context.Background(), query, req.Title, req.Content, claims.UserID).Scan(&newID)
	if err != nil {
		http.Error(w, "Failed to create announcement", http.StatusInternalServerError)
		return
	}

	response := models.Announcement{
		ID:       newID,
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: claims.UserID,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
func GetAnnouncementHandler(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, title, content, author_id, created_at FROM announcements ORDER BY created_at DESC`
	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		http.Error(w, "Failed to fetch announcements", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var announcement []models.Announcement
	for rows.Next() {
		var a models.Announcement
		err := rows.Scan(&a.ID, &a.Title, &a.Content, &a.AuthorID, &a.CreatedAt)
		if err != nil {
			http.Error(w, "failed to read announcement data", http.StatusInternalServerError)
			return
		}
		announcement = append(announcement, a)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(announcement)

}
