package cmd

import (
	"fmt"
	"log"
	"moco/config"
	"net/http"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop [id]",
	Short: "Stop time tracking for a given project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("args", args)
		config := config.Init()
		apiKey := config.GetString("api_key")

        if apiKey == "" {
            log.Fatal("api_key is required")
        }

        client := &http.Client{}

        url := fmt.Sprintf("https://foobaragency.mocoapp.com/api/v1/projects/%s/start", args[0])
        req, err := http.NewRequest("POST", url, nil)
        resp, err := client.Do(req)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(resp.StatusCode)

	},
}
