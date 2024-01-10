package data

import (
	"encoding/json"
	"fmt"
	"io"
	"moco/config"
	"net/http"
)

type User struct {
    Id   int    `json:"id"`
    FirstName string `json:"firstname"`
    LastName string `json:"lastname"`
}

func GetUserId() (int, error) {
	config := config.Init()

	apiKey := config.GetString("api_key")
    firstName := config.GetString("first_name")
    lastName := config.GetString("last_name")
    domain := config.GetString("domain")
	if apiKey == "" {
		return 0, fmt.Errorf("api_key not set")
	} else if firstName == "" {
        return 0, fmt.Errorf("first_name not set")
	} else if lastName == "" {
        return 0, fmt.Errorf("last_name not set")
	} else if domain == "" {
        return 0, fmt.Errorf("domain not set")
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf("https://%s.mocoapp.com/api/v1/users", domain), nil)
    req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", apiKey))

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return 0, fmt.Errorf("Could not find user")
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return 0, fmt.Errorf("Could not find user")
    }

    var users []User

    json.Unmarshal(body, &users)
    for _, user := range users {
        if user.FirstName == firstName && user.LastName == lastName {
            return user.Id, nil
        }
    }
    return 0, fmt.Errorf("Could not find user")
}
