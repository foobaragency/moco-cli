package cmd

import (
	"fmt"
	"io"
	"log"
	"moco/config"
	"net/http"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start time tracking for a given project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config := config.Init()
		apiKey := config.GetString("api_key")

		req, _ := http.NewRequest("GET", "https://foobaragency.mocoapp.com/api/v1/projects/assigned", nil)
		req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", apiKey))
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERROR] -", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error while reading the response bytes:", err)
		}
		log.Println(string([]byte(body)))
	},
}
