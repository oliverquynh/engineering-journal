package main

import (
	"fmt"
	"log"
	"net/http"
)

// defaultHandler handles requests to the "/" route
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	// Only respond to GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Engineering Journal")
}

// helloHandler handles requests to the "/hello" route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Only respond to GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	// 1. Create a new request multiplexer (router)
	mux := http.NewServeMux()

	// 2. Register routes and their matching handler functions
	mux.HandleFunc("/", defaultHandler)
	mux.HandleFunc("/hello", helloHandler)

	// 3. Configure the server settings
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// 4. Start listening for incoming requests
	fmt.Println("Server is running on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
