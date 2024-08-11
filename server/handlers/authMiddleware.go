package handlers

import (
	"auth/database"
	"context"
	"database/sql"
	"net/http"
	"strings"
)

type ctxKey string
const userID ctxKey = "userID"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r * http.Request) {
		authString := r.Header.Get("Authorization")
		if authString == "" {
			sendErrorResponse(w, "Bearer token required", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authString, "Bearer ")
		if token == authString {
			sendErrorResponse(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		var id int
		err := database.DB.QueryRow("SELECT id FROM test_users WHERE token = $1", token).Scan(&id)
		if err == sql.ErrNoRows {
			sendErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if err != nil {
			sendErrorResponse(w, "Server error", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), userID, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}