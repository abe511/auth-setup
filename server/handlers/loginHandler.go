package handlers

import (
	"auth/database"
	"auth/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)


func sendErrorResponse(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	errorResponse := map[string]string{"message": message}
	json.NewEncoder(w).Encode(errorResponse)
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)
	return re.MatchString(email)
}

// password must contain at least 1 uppercase, 1 lowercase letter, 1 digit, 1 special character
// password minimum length is 8
func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	uppercase := regexp.MustCompile(`[A-Z]`)
	if !uppercase.MatchString(password) {
		return false
	}

	lowercase := regexp.MustCompile(`[a-z]`)
	if !lowercase.MatchString(password) {
		return false
	}

	digit := regexp.MustCompile(`[0-9]`)
	if !digit.MatchString(password) {
		return false
	}

	special := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)

	return special.MatchString(password)
}

func comparePasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	fmt.Printf("email: %v, password: %v\n", user.Email, user.Password)
	
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	// validate request
	if user.Email == "" || user.Password == "" {
		sendErrorResponse(w, "User email and password are required", http.StatusBadRequest)
		return
	}

	if !isValidEmail(user.Email) {
		sendErrorResponse(w, "Invalid email", http.StatusBadRequest)
		return
	}

	if !isValidPassword(user.Password) {
		sendErrorResponse(w, "Invalid password", http.StatusBadRequest)
		return
	}

	fmt.Println("credentials checked")
	
	// find the user by email in the db
	var id int
	var passwordHash string
	err = database.DB.QueryRow("SELECT id, password FROM test_users WHERE email = $1", user.Email).Scan(&id, &passwordHash)
	// if not found - respond with error 403
	if err == sql.ErrNoRows {
		sendErrorResponse(w, "User not found", http.StatusForbidden)
		return
	}
	if err != nil {
		sendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// if found - check password hashes
	if !comparePasswordHash(passwordHash, user.Password) {
		sendErrorResponse(w, "Wrong password", http.StatusForbidden)
		return
	}
	
	// if password hashes match - generate token and send it with response status 200
	token := fmt.Sprintf("Bearer %d", time.Now().Unix())
	w.Header().Set("Content-type", "application/json")
	response := map[string]string{"accessToken": token}
	json.NewEncoder(w).Encode(response)
}