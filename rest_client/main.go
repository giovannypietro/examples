// Example to call a mock user API and search for a user with a specific ID.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
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

	// Read the response body
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

func main() {

	var id int
	flag.IntVar(&id, "id", 0, "an user's id")
	flag.Parse()

	if flag.NFlag() == 0 {
		log.Fatalf("Please provide an integer using the -id flag.")
	}

	user, err := GetUser(id)

	if err != nil {
		log.Fatalf("Error getting user: %s", err)
	}

	// Print JSON contents
	fmt.Printf("ID: %d\n", user.ID)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("Phone: %s\n", user.Phone)
	fmt.Printf("Company: %s\n", user.Company.Name)
	fmt.Printf("CatchPhrase: %s\n", user.Company.CatchPhrase)
	fmt.Printf("BS: %s\n", user.Company.Bs)
}
