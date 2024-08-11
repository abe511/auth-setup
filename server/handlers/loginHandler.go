package handlers

import (
	"auth/database"
	"auth/models"
	"auth/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	errorResponse := map[string]string{"message": message}
	json.NewEncoder(w).Encode(errorResponse)
}

// validate request
func hasValidCredentials(w http.ResponseWriter ,user models.User) bool {
	if user.Email == "" || user.Password == "" {
		sendErrorResponse(w, "User email and password are required", http.StatusBadRequest)
		return false
	}
	if !utils.IsValidEmail(user.Email) {
		sendErrorResponse(w, "Invalid email", http.StatusBadRequest)
		return false
	}
	if !utils.IsValidPassword(user.Password) {
		sendErrorResponse(w, "Invalid password", http.StatusBadRequest)
		return false
	}
	return true
}

// search user by email in the db
func searchByEmail(w http.ResponseWriter, user models.User) (string, error) {
	var passwordHash string
	err := database.DB.QueryRow("SELECT password FROM test_users WHERE email = $1", user.Email).Scan(&passwordHash)
	// if user not found - respond with 403 error
	if err != nil {
		if err == sql.ErrNoRows {
			sendErrorResponse(w, "User not found", http.StatusForbidden)
		} else {
			sendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		}
	}
	return passwordHash, err
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)	
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !hasValidCredentials(w, user) {
		return
	}

	passwordHash, err := searchByEmail(w, user)
	if err != nil {
		return
	}
	// if found - check password hashes
	if !utils.ComparePasswordHashes(passwordHash, user.Password) {
		sendErrorResponse(w, "Wrong password", http.StatusForbidden)
		return
	}
	// if password hashes match - generate token
	token, err := utils.HashString(user.Email + user.Password)
	if err != nil {
		sendErrorResponse(w, "Token generation failed", http.StatusInternalServerError)
	}
	// write new token to the db
	_, err = database.DB.Exec("UPDATE test_users SET token = $1 WHERE email = $2 ", token, user.Email)
	if err != nil {
		sendErrorResponse(w, "Token update failed", http.StatusInternalServerError)
		return
	}
	// send the token with response status 200
	tokenString := fmt.Sprintf("Bearer %s", token)
	w.Header().Set("Content-type", "application/json")
	response := map[string]string{"accessToken": tokenString}
	json.NewEncoder(w).Encode(response)
}
