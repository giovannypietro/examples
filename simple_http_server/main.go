package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Request struct {
	Request any
}

type Message struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/echo", echoHandler)

	log.Println("Starting server on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// pingHandler responds with a "pong!" message
func pingHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"message": "pong!",
	}

	if err := writeJSONResponse(w, http.StatusOK, data); err != nil {
		log.Printf("could not write response: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// echoHandler echoes back the received JSON payload
func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only service.", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read the request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var jsonBody map[string]interface{}
	if err := json.Unmarshal(body, &jsonBody); err != nil {
		http.Error(w, "Invalid JSON. Please check your JSON document.", http.StatusBadRequest)
		return
	}

	if err := writeJSONResponse(w, http.StatusOK, jsonBody); err != nil {
		log.Printf("could not write response: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// writeJSONResponse writes a JSON response to the ResponseWriter
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("could not marshal json: %w", err)
	}

	_, err = w.Write(jsonData)
	return err
}
