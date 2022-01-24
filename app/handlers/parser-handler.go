package handlers

import (
	"encoding/json"
	"net/http"
)
type request struct {
	Data string
}

type errorResponse struct {
	Error string `json:"error"`
}

type successResponse struct {
	Occurences []wordOccurence `json:"occurences"`
}

type wordOccurence struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}
// Function to return bad request to response
func returnBadRequest(message string, w http.ResponseWriter) {
	// Set content type header
	w.Header().Set("Content-Type", "application/json")
	// Set response status
	w.WriteHeader(http.StatusBadRequest)
	response := errorResponse{message}

	json.NewEncoder(w).Encode(response)
}

// Function to return successful response
func returnSuccessfulResponse(message []wordOccurence, w http.ResponseWriter) {
	// Set content type headaer
	w.Header().Set("Content-Type", "application/json")
	// Set response status
	w.WriteHeader(http.StatusOK)
	response := successResponse{message}

	json.NewEncoder(w).Encode(response)
}
