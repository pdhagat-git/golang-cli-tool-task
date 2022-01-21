package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error: Unable to load env")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Welcome to GoLang text parser")
	}).Methods("GET")

	log.Printf("Server listening on port: %s", os.Getenv("PORT"))
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}
