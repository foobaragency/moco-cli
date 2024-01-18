package cmd

import (
	"fmt"
	"moco/data"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var stopCmd = &cobra.Command{
	Use:   "stop [id]",
	Short: "Stop time tracking for all projects",
	Run: func(cmd *cobra.Command, args []string) {
        activities, err := data.GetActivities()
        if err != nil {
            fmt.Println(err)
            return
        }
        for _, activity := range activities {
            if activity.TimerStartedAt != "" {
                data.StopActivity(activity.Id)
            }
        }
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
