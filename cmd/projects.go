package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"moco/config"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

type Task struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
    Active bool `json:"active"`
    Billable bool `json:"billable"`
}

type Project struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
    Active bool `json:"active"`
    Tasks []Task `json:"tasks"`
}

// rootCmd represents the base command when called without any subcommands
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "List projects assigned to you",
	Run: func(cmd *cobra.Command, args []string) {
        config := config.Init()

		apiKey := config.GetString("api_key")
		if apiKey == "" {
			fmt.Println("api_key not set")
			return
		}

		req, _ := http.NewRequest("GET", "https://foobaragency.mocoapp.com/api/v1/projects/assigned", nil)
		req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", apiKey))
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERROR] -", err)
		}
		defer resp.Body.Close()

        var projects []Project

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error while reading the response bytes:", err)
		}
        json.Unmarshal(body, &projects)

        var projectNames []string
        for _, project := range projects {
            projectNames = append(projectNames, project.Name)
        }

        fmt.Println(strings.Join(projectNames, "\n"))
	},
}
