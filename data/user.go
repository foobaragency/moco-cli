package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"moco/config"
	"net/http"
)

type User struct {
    Id   int    `json:"id"`
    FirstName string `json:"firstname"`
    LastName string `json:"lastname"`
}

func GetUserId() int {
	config := config.Init()

	apiKey := config.GetString("api_key")
    firstName := config.GetString("first_name")
    lastName := config.GetString("last_name")
	if apiKey == "" {
		log.Fatal("api_key not set")
	}

	req, _ := http.NewRequest("GET", "https://foobaragency.mocoapp.com/api/v1/users", nil)
    req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", apiKey))

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("Error on response.\n[ERROR] -", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal("Error while reading the response bytes:", err)
    }

    var users []User

    json.Unmarshal(body, &users)
    for _, user := range users {
        if user.FirstName == firstName && user.LastName == lastName {
            return user.Id
        }
    }
    return 0
}
