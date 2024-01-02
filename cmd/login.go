package cmd

import (
	"bufio"
	"fmt"
	"os"

	// "io"
	// "log"
	"moco/config"

	"github.com/spf13/cobra"
)

type User struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// rootCmd represents the base command when called without any subcommands
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to Moco",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.Init()

		apiKey := config.GetString("api_key")
		if apiKey != "" {
			fmt.Println("Looks like you're already logged in.")
			return
		}
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter your first name: ")
		scanner.Scan()
		firstName := scanner.Text()
		fmt.Print("Enter your last name: ")
		scanner.Scan()
		lastName := scanner.Text()
		fmt.Print("Your API key: ")
		scanner.Scan()
		apiKey = scanner.Text()

		config.Set("api_key", apiKey)
		config.Set("first_name", firstName)
		config.Set("last_name", lastName)

        config.WriteConfig()
	},
}
