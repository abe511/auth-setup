package handlers

import (
	"auth/database"
	"auth/models"
	"auth/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)	
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	// validate request
	if user.Email == "" || user.Password == "" {
		sendErrorResponse(w, "User email and password are required", http.StatusBadRequest)
		return
	}
	if !utils.IsValidEmail(user.Email) {
		sendErrorResponse(w, "Invalid email", http.StatusBadRequest)
		return
	}
	if !utils.IsValidPassword(user.Password) {
		sendErrorResponse(w, "Invalid password", http.StatusBadRequest)
		return
	}	
	// search by email in the db
	var id int
	var passwordHash string
	err = database.DB.QueryRow("SELECT id, password FROM test_users WHERE email = $1", user.Email).Scan(&id, &passwordHash)
	// if user not found - respond with error 403
	if err == sql.ErrNoRows {
		sendErrorResponse(w, "User not found", http.StatusForbidden)
		return
	}
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// if found - check password hashes
	if !utils.ComparePasswordHashes(passwordHash, user.Password) {
		sendErrorResponse(w, "Wrong password", http.StatusForbidden)
		return
	}
	// if password hashes match - generate token and send it with response status 200
	token := fmt.Sprintf("Bearer %d", time.Now().Unix())
	w.Header().Set("Content-type", "application/json")
	response := map[string]string{"accessToken": token}
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	errorResponse := map[string]string{"message": message}
	json.NewEncoder(w).Encode(errorResponse)
}