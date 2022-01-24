package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

// Load environment variables
func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error: Unable to load env: ", err)
	}
}

// Function to call API service to get top 10 most occured words in text string
func getTopWords(text string) string {
	// Create request body
	requestBody := map[string]string{"data": text}
	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal("Error: Unable to create request body: ", err)
	}

	// Call api to get top 10 most occurred words
	response, err := http.Post(os.Getenv("API_URL"), "application/json", bytes.NewBuffer(requestJSON))
	if err != nil {
		log.Fatal("Error: Failed API call: ", err)
	}

	defer response.Body.Close()

	// Read response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error: Unable to read API response body: ", err)
	}

	// Prettify json response
	var topWords bytes.Buffer
	err = json.Indent(&topWords, responseBody, "", "  ")
	if err != nil {
		log.Fatal("Error: Unable to parse response body: ", err)
	}

	return topWords.String()
}

// Function to read file content at given path
func readFile(filePath string) string {
	extension := filepath.Ext(filePath)
	if extension != ".txt" {
		log.Fatal("Error: Please provide a .txt file path")
	}

	// Open file at filePath
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error: Unable to open file: ", err)
	}

	defer file.Close()

	// Read contents of the file
	text, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Error: Unable to read file: ", err)
	}

	return string(text)
}

func main() {
	var filePath string

	// Initialize cli app and define flags
	app := &cli.App{
		Name:  "Parser",
		Usage: "Parse files and return word count",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "path",
				Usage:       "path to text file",
				EnvVars:     []string{"FILE_PATH"}, // Default file path set from env
				Destination: &filePath,
			},
		},
		Action: func(c *cli.Context) error {
			text := readFile(filePath)
			topWords := getTopWords(text)
			log.Println(topWords)
			return nil
		},
	}

	// Run cli app
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("Error: Unable to start app: ", err)
	}
}
