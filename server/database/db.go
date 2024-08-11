package database

import (
	"auth/models"
	"auth/utils"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")

	port, err := strconv.Atoi(portStr)
	utils.LogFatalError("Invalid port number: ", err)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, sslMode)

	DB, err = sql.Open("postgres", connStr)
	utils.LogFatalError("Database connection failed: ", err)

	err = DB.Ping()
	utils.LogFatalError("Database ping failed: ", err)

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS test_users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		token VARCHAR(255) NOT NULL
		)
	`)
	utils.LogFatalError("Users table creation failed: ", err)	

	populateDB()
}

func populateDB() {
	users := []models.User{
		{Email: "user1@server.net", Password: "Pa$$word1"},
		{Email: "user2@server.net", Password: "Pa$$word2"},
		{Email: "user3@server.net", Password: "Pa$$word3"},
	}
	
	for _, user := range users {
		var err error
		hashedPassword, err := utils.HashString(user.Password)
		if err != nil {
			log.Printf("Password hashing failed for %s: %v", user.Email, err)
			continue
		}

		var id int
		token := ""
		// insert the user only if it does not exist
		err = DB.QueryRow(`
			INSERT INTO test_users (email, password, token)
			VALUES ($1, $2, $3)
			ON CONFLICT (email) DO NOTHING
			RETURNING id
		`, user.Email, hashedPassword, token).Scan(&id)

		if err == sql.ErrNoRows {
			fmt.Printf("User already exists: %s\n", user.Email)
		} else if err != nil {
			log.Printf("User insertion failed for %s: %v", user.Email, err)
		} else {
			fmt.Printf("New user added. id: %d, email: %s\n", id, user.Email)
		}
	}

	fmt.Println("Database population completed")
}