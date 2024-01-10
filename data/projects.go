package data

import (
	"encoding/json"
	"fmt"
	"io"
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

func GetProject(projectId int) (Project, error) {
    config := config.Init()
    apiKey := config.GetString("api_key")
    if apiKey == "" {
        fmt.Errorf("api_key not set")
    }
    
    req, _ := http.NewRequest("GET", "https://foobaragency.mocoapp.com/api/v1/projects/assigned", nil)
    req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", apiKey))
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return Project{}, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return Project{}, err
    }

    var project []Project
    json.Unmarshal(body, &project)
    for _, p := range project {
        if p.Id == projectId {
            return p, nil
        }
    }
    return Project{}, fmt.Errorf("project not found")
}

func GetProjects() ([]Project, error) {
    config := config.Init()
    apiKey := config.GetString("api_key")
    if apiKey == "" {
        return []Project{}, fmt.Errorf("api_key not set")
    }
    
    req, _ := http.NewRequest("GET", "https://foobaragency.mocoapp.com/api/v1/projects/assigned", nil)
    req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", apiKey))
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return []Project{}, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return []Project{}, err
    }

    var projects []Project
    json.Unmarshal(body, &projects)
    return projects, nil
}
