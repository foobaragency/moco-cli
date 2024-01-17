package cmd

import (
	"fmt"
	"moco/data"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var stopCmd = &cobra.Command{
	Use:   "stop [id]",
	Short: "Stop time tracking for a given project",
	Run: func(cmd *cobra.Command, args []string) {
        if activityId != 0 {
            data.StopActivity(activityId)
        }
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
    stopCmd.Flags().IntVarP(&activityId, "activity", "a", 0, "Activity ID (if not provided, a new activity will be created)")
    stopCmd.MarkFlagRequired("activity")
	rootCmd.AddCommand(stopCmd)
}
