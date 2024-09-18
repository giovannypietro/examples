package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/strata-io/service-extension/orchestrator"
)

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type User struct {
	ID      int     `json:"id"`
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Company Company `json:"company"`
}

func GetUser(api orchestrator.Orchestrator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		client := &http.Client{}

		url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", 2)

		resp, err := client.Get(url)

		if err != nil {
			http.Error(w, "Failed to obtain token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			http.Error(w, "Failed to read response body: "+err.Error(), http.StatusInternalServerError)
			return
		}

		var user User

		if err := json.Unmarshal(body, &user); err != nil {
			http.Error(w, "Failed to unmarshall identity: "+err.Error(), http.StatusInternalServerError)
			return
		}

		encodeJson(w, user)
	}
}

func encodeJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
func Serve(api orchestrator.Orchestrator) error {

	var (
		logger = api.Logger()
		router = api.Router()
	)

	logger.Info("se", "exposing Identity API")

	err := router.HandleFunc("/identity", GetUser(api))

	if err != nil {
		return fmt.Errorf("failed to handle route: %w", err)
	}

	return nil
}
