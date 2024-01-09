package cmd

import (
	"fmt"
	"os"

	// "io"
	// "log"
	"moco/config"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out from Moco",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.Init()

		apiKey := config.GetString("api_key")
		if apiKey == "" {
			fmt.Println("Looks like you're already logged out.")
			return
		}
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			os.Exit(1)
		}
		os.Remove(config.ConfigFileUsed())
		os.Remove(home + "/.config/moco")
	},
}
