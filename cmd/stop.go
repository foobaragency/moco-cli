package cmd

import (
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
	},
}

func init() {
	stopCmd.Flags().IntVarP(&projectId, "project", "p", 0, "Project ID")
	stopCmd.Flags().IntVarP(&taskId, "task", "t", 0, "Task ID")
    stopCmd.Flags().IntVarP(&activityId, "activity", "a", 0, "Activity ID (if not provided, a new activity will be created)")
	rootCmd.AddCommand(stopCmd)
}
