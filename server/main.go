package main

import (
	"fmt"
	"log"
	"net/http"
	"auth/handlers"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("POST /login", handlers.LoginHandler)
	router.HandleFunc("GET /profile", handlers.GetProfile)

	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	fmt.Printf("Server running on port: %v\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
