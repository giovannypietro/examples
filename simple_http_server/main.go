// Very simple implementation of an HTTP server in Go.

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

	http.HandleFunc("/ping", pinghandler)

	http.HandleFunc("/echo", echohandler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func pinghandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"message": "pong!",
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s", string(jsonData))
}

func echohandler(w http.ResponseWriter, r *http.Request) {

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s", body)
}
