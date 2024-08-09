package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type User struct {
	Email 		string	`json:"email"`
	Password	string	`json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Printf("email: %v, password: %v\n", user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// search the user by username in the db
	// if not found - respond with error
	// if found - check password hashes
	// if password hashes match - generate token and send it with the response

}