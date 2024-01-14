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

type Pickable interface {
    GetName() string
    GetId() int
    FilterValue() string
}

func (p Project) FilterValue() string { return p.Name }
func (p Project) GetName() string    { return p.Name }
func (p Project) GetId() int      { return p.Id }

func (t Task) FilterValue() string { return t.Name }
func (t Task) GetName() string    { return t.Name }
func (t Task) GetId() int      { return t.Id }

func GetProject(projectId int) (Project, error) {
    config := config.Init()
    apiKey := config.GetString("api_key")
    domain := config.GetString("domain")
    if apiKey == "" {
        return Project{}, fmt.Errorf("api_key not set")
    }
    if domain == "" {
        return Project{}, fmt.Errorf("domain not set")
    }

    
    req, _ := http.NewRequest("GET", fmt.Sprintf("https://%s.mocoapp.com/api/v1/projects/assigned", domain), nil)
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
    domain := config.GetString("domain")
    if apiKey == "" {
        return []Project{}, fmt.Errorf("api_key not set")
    }
    if domain == "" {
        return []Project{}, fmt.Errorf("domain not set")
    }
    
    req, _ := http.NewRequest("GET", fmt.Sprintf("https://%s.mocoapp.com/api/v1/projects/assigned", domain), nil)
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
