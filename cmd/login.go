package cmd

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"moco/config"

	"github.com/spf13/cobra"
)

type User struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to Moco",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.Init()

		apiKey := config.GetString("api_key")
		if apiKey != "" && config.GetString("first_name") != "" && config.GetString("last_name") != "" && config.GetString("domain") != "" {
			fmt.Println("Looks like you're already logged in.")
			return
		}
		var domain string
		var firstName string
		var lastName string

		form := huh.NewForm(
            huh.NewGroup(
                huh.NewInput().Title("Domain").Prompt("?").Value(&domain),
                huh.NewInput().Title("First Name").Prompt("?").Value(&firstName),
                huh.NewInput().Title("Last Name").Prompt("?").Value(&lastName),
                huh.NewInput().Title("API Key").Prompt("?").Value(&apiKey).Password(true),
            ),
        )

        form.Run()

		config.Set("api_key", apiKey)
		config.Set("first_name", firstName)
		config.Set("last_name", lastName)
		config.Set("domain", domain)

		config.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
