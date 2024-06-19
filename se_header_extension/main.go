// Example Maverics service extension that fetches a user from a mock API using an id and returns the email address
// to the orchestrator.

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

func GetUser(userID int) (*User, error) {

	client := &http.Client{}

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", userID)

	resp, err := client.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var user User

	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateEmailHeader(api orchestrator.Orchestrator, _ http.ResponseWriter, _ *http.Request) (http.Header, error) {

	logger := api.Logger()

	logger.Info("Service Extension", "Building email custom header from remote API.")

	session, err := api.Session()

	if err != nil {
		return nil, fmt.Errorf("unable to retrieve session: %w", err)
	}

	// Replace with any attribute available via the Orchestrator session.

	id, err := session.GetInt("azure.id")

	if err != nil {
		return nil, fmt.Errorf("unable to retrieve the attribute 'azure.id from the orchestrator': %w", err)
	}

	user, err := GetUser(id)

	if err != nil {
		return nil, fmt.Errorf("unable to fetch the user from the remote API: %w", err)
	}

	email := user.Email

	// If all ok make the custom header and return to the Orchestrator.

	header := make(http.Header)

	header["CUSTOM-EMAIL"] = []string{email}

	return header, nil
}
