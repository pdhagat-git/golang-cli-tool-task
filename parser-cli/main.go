package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Load environment variables
func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error: Unable to load env: ", err)
	}
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

	//Print contents of file
	data := readFile(os.Getenv("FILE_PATH"))
	log.Println(data)

}
