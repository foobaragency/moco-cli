package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"moco/config"
	"net/http"
)

type Task struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

type Project struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
    Tasks []Task
}


func GetProjects() []Project {
    config := config.Init()
    apiKey := config.GetString("api_key")
    if apiKey == "" {
        log.Fatal("api_key not set")
    }
    
    req, _ := http.NewRequest("GET", "https://foobaragency.mocoapp.com/api/v1/projects/assigned", nil)
    req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", apiKey))
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal("Error on response.\n[ERROR] -", err)
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Println("Error while reading the response bytes:", err)
    }

    var projects []Project
    json.Unmarshal(body, &projects)
    return projects
}
