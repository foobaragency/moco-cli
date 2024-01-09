package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"moco/config"
	"net/http"
)

type Activity struct {
	Id             int     `json:"id"`
	Date           string  `json:"date"`
	Hours          float64 `json:"hours"`
	Seconds        int     `json:"seconds"`
	Description    string  `json:"description"`
	TimerStartedAt string  `json:"timer_started_at"`
}

func GetActivities() []Activity {
	config := config.Init()

	apiKey := config.GetString("api_key")
	if apiKey == "" {
		log.Fatal("api_key not set")
	}

	req, _ := http.NewRequest("GET", "https://foobaragency.mocoapp.com/api/v1/activites", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", apiKey))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	var activites []Activity

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error while reading the response bytes:", err)
	}
	json.Unmarshal(body, &activites)

	return activites
}
