package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define a struct to map the JSON data structure (change this based on the API you are using)
type ApiResponse struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Handler function to fetch data from an API and render HTML
func handler(w http.ResponseWriter, r *http.Request) {
	// Fetch data from an external API (for example, JSONPlaceholder API)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		http.Error(w, "Error fetching data from API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decode the JSON response into a Go struct
	var data ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, "Error decoding JSON data", http.StatusInternalServerError)
		return
	}

	// Render HTML and display the API data
	fmt.Fprintf(w, "<html><body>")
	fmt.Fprintf(w, "<h1>API Data:</h1>")
	fmt.Fprintf(w, "<p>User ID: %d</p>", data.UserID)
	fmt.Fprintf(w, "<p>Task ID: %d</p>", data.ID)
	fmt.Fprintf(w, "<p>Title: %s</p>", data.Title)
	fmt.Fprintf(w, "<p>Completed: %v</p>", data.Completed)
	fmt.Fprintf(w, "</body></html>")
}

func main() {
	http.HandleFunc("/", handler)

	// Listen on a different port (9000)
	fmt.Println("Server is listening on port 9000...")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
