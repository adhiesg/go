package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

// Global variable to store the name
var myName = "John Doe"
var mutex sync.Mutex

// Handler function for the /api endpoint (GET request)
func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Create the response message with the predefined variable
	response := map[string]string{"message": "Hello, this is " + myName}

	// Convert the response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonResponse)
}

// Handler function for the /setname endpoint (POST request)
func setNameHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body
	var requestData map[string]string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Extract the new name from the request body
	newName, ok := requestData["name"]
	if !ok {
		http.Error(w, "Missing 'name' field in JSON body", http.StatusBadRequest)
		return
	}

	// Update the global variable with the new name
	mutex.Lock()
	myName = newName
	mutex.Unlock()

	// Send a success response
	response := map[string]string{"message": "Name updated successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonResponse)
}

func main() {
	// Define the endpoint and the corresponding handler function for GET request
	http.HandleFunc("/api", apiHandler)

	// Define the endpoint and the corresponding handler function for POST request
	http.HandleFunc("/setname", setNameHandler)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
