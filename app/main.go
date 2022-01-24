package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"parser-service/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Load environment variables
func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error: Unable to load env")
	}
}

func main() {
	// Initialize mux router
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Welcome to GoLang text parser")
	}).Methods("GET")

	// Define paths and call to handlers package
	router.HandleFunc("/api/v1/parser/wordCount", handlers.WordCount).Methods("POST")
	log.Printf("Server listening on port: %s", os.Getenv("PORT"))
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}
