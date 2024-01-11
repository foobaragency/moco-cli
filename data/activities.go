package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"moco/config"
	"net/http"
	"time"
)

type Activity struct {
	Id             int     `json:"id"`
	Date           string  `json:"date"`
	Hours          float64 `json:"hours"`
	Seconds        int     `json:"seconds"`
	Description    string  `json:"description"`
	TimerStartedAt string  `json:"timer_started_at"`
	Project        struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	Task struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	User struct {
		Id        int    `json:"id"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
	}
}

func StartActivity(activityId int) error {
	config := config.Init()
	apiKey := config.GetString("api_key")
	if apiKey == "" {
        return fmt.Errorf("api_key not set")
	}
	domain := config.GetString("domain")
	if domain == "" {
        return fmt.Errorf("domain not set")
	}

	req, _ := http.NewRequest("PATCH", fmt.Sprintf("https://%s.mocoapp.com/api/v1/activities/%d/start_timer", domain, activityId), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", apiKey))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode == 422 {
        return fmt.Errorf("Something went wrong")
	} else if resp.StatusCode == 404 {
        return fmt.Errorf("Activity not found")
    }
	defer resp.Body.Close()
    return nil
}

func StopActivity(activityId int) error {
	config := config.Init()
	apiKey := config.GetString("api_key")
	if apiKey == "" {
        return fmt.Errorf("api_key not set")
	}
	domain := config.GetString("domain")
	if domain == "" {
        return fmt.Errorf("domain not set")
	}

	req, _ := http.NewRequest("PATCH", fmt.Sprintf("https://%s.mocoapp.com/api/v1/activities/%d/stop_timer",domain, activityId), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", apiKey))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode == 422 {
        return fmt.Errorf("Error on response.\n[ERROR] - %s", err)
	}
	defer resp.Body.Close()
    return nil
}

func CreateActivity(projectId int, taskId int, description string) error {
	config := config.Init()
	apiKey := config.GetString("api_key")
    domain := config.GetString("domain")
	if apiKey == "" {
        return fmt.Errorf("api_key not set")
	}
    if domain == "" {
        return fmt.Errorf("domain not set")
    }

	type ActivityBody struct {
		ProjectId   int    `json:"project_id"`
		TaskId      int    `json:"task_id"`
		Description string `json:"description"`
		Date        string `json:"date"`
	}

    body := ActivityBody{
        ProjectId:   projectId,
        TaskId:      taskId,
        Description: description,
        Date:        time.Now().Format("2006-01-02"),
    }
    marshaledBody, err := json.Marshal(body)
    if err != nil {
        log.Fatal(err)
    }

	req, _ := http.NewRequest("POST", fmt.Sprintf("https://%s.mocoapp.com/api/v1/activities", domain), bytes.NewReader(marshaledBody))
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", apiKey))
    req.Header.Add("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil || resp.StatusCode == 422 {
        return err
    }
    defer resp.Body.Close()
    return nil
}

func DeleteActivity(id int) error {
    config := config.Init()
    apiKey := config.GetString("api_key")
    domain := config.GetString("domain")
    if apiKey == "" {
        return fmt.Errorf("api_key not set")
    }
    if domain == "" {
        return fmt.Errorf("domain not set")
    }

    req, _ := http.NewRequest("DELETE", fmt.Sprintf("https://%s.mocoapp.com/api/v1/activities/%d", domain, id), nil)
    req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", apiKey))
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil || resp.StatusCode >= 400 && resp.StatusCode <= 499 {
        return err
    }
    defer resp.Body.Close()
    return nil
}

func GetActivities() ([]Activity, error) {
	config := config.Init()

	apiKey := config.GetString("api_key")
    domain := config.GetString("domain")
	if apiKey == "" {
        return nil, fmt.Errorf("api_key not set")
	}
    if domain == "" {
        return nil, fmt.Errorf("domain not set")
    }

    userId, err := GetUserId()
    if err != nil {
        return nil, err
    }
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://%s.mocoapp.com/api/v1/activities?user_id=%d", domain, userId), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", apiKey))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	var activities []Activity

	body, err := io.ReadAll(resp.Body)
	if err != nil {
        return nil, err
	}
	json.Unmarshal(body, &activities)

	return activities, nil
}
