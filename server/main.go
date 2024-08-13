package main

import (
	"auth/handlers"
	"auth/database"
	"auth/utils"
	"fmt"
	"log"
	"net/http"
)


func main() {

	database.InitDB()

	router := http.NewServeMux()

	router.HandleFunc("POST /login", handlers.LoginHandler)
	router.HandleFunc("GET /profile", handlers.AuthMiddleware(handlers.GetProfile))
	
	corsEnabledRouter := utils.GlobalCORS(router)

	server := http.Server{
		Addr: ":8080",
		Handler: corsEnabledRouter,
	}

	fmt.Printf("Server running on port: %v\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
