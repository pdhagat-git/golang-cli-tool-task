package handlers

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"
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

// Function to handle word count request
func WordCount(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req request

	// Decode request body
	err := decoder.Decode(&req)
	if err != nil {
		returnBadRequest("Invalid request body", w)
		return
	}

	if req.Data == "" {
		returnBadRequest("Invalid request body", w)
		return
	}

	// Get word counts
	wordCount := getWordCount(req.Data)

	returnSuccessfulResponse(wordCount, w)
}

// Function to handle top words count request
func TopWordsCount(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req request

	// Decode request body
	err := decoder.Decode(&req)
	if err != nil {
		returnBadRequest("Invalid request body", w)
		return
	}

	if req.Data == "" {
		returnBadRequest("Invalid request body", w)
		return
	}

	// Get top 10 most occured words in text
	topWordsCount := getTopWordsCount(req.Data)

	returnSuccessfulResponse(topWordsCount, w)
}

// Function to get occurences of all words in text string
func getWordCount(text string) (wordOccurences []wordOccurence) {
	// Convert string text to slice array by separating through white spaces
	words := strings.Fields(text)

	// Map words to occurences in text
	wordsMap := make(map[string]int)
	for _, word := range words {
		wordsMap[word]++
	}

	// Append word occurences tp struct
	for word, count := range wordsMap {
		wordOccurences = append(wordOccurences, wordOccurence{word, count})
	}

	return
}

// Function to get top 10 word occurences in text string
func getTopWordsCount(text string) (wordOccurences []wordOccurence) {
	// Get occurenced of all words in text
	wordOccurences = getWordCount(text)

	// Sort in decreasing order of word counts
	sort.Slice(wordOccurences, func(i, j int) bool {
		return wordOccurences[i].Count > wordOccurences[j].Count
	})

	// Get the first 10 elements of slice
	if len(wordOccurences) > 10 {
		wordOccurences = wordOccurences[:10]
	}

	return wordOccurences
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
