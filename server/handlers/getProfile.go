package handlers

import (
	"auth/database"
	"auth/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value(userID).(int)
	if !ok {
		sendErrorResponse(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	var email string
	err := database.DB.QueryRow("SELECT (email) FROM test_users WHERE id = $1", id).Scan(&email)
	if err == sql.ErrNoRows {
		sendErrorResponse(w, "Wrong user id", http.StatusUnauthorized)
		return
	}
	if err != nil {
		sendErrorResponse(w, "Database query failed", http.StatusInternalServerError)
		return
	}

	index := strings.Index(email, "@");
	if  index == -1 {
		sendErrorResponse(w, "Invalid email", http.StatusInternalServerError)
	}
	username := email[:index]

	w.Header().Set("Content-type", "application/json")
	response := models.UserProfile{Email: email, Username: username}
	json.NewEncoder(w).Encode(response)
}