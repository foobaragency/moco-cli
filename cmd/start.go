package cmd

import (
	"fmt"
	"moco/data"

	"github.com/spf13/cobra"
)

var activityId int
var projectId int
var description string
var taskId int

// rootCmd represents the base command when called without any subcommands
var startCmd = &cobra.Command{
	Use:   "start [id]",
	Short: "Start time tracking",
	Long:  "Start time for an activity. You may either provide:\n1. An activity ID\n2. A project ID, task ID, and a description",
	Run: func(cmd *cobra.Command, args []string) {
		if activityId != 0 {
			err := data.StartActivity(activityId)
            if err != nil {
                fmt.Printf("Error starting activity: %v\n", err)
            }
		} else if projectId != 0 && taskId != 0 && description != "" {
			data.CreateActivity(projectId, taskId, description)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	startCmd.Flags().IntVarP(&projectId, "project", "p", 0, "Project ID")
	startCmd.Flags().IntVarP(&taskId, "task", "t", 0, "Task ID")
	startCmd.Flags().StringVarP(&description, "description", "d", "", "Description")
	startCmd.Flags().IntVarP(&activityId, "activity", "a", 0, "Activity ID")
	startCmd.SetUsageTemplate(`Usage 
  moco start -a ACTIVITY
  moco start -p PROJECT -t TASK -d DESCRIPTION
`)
	rootCmd.AddCommand(startCmd)
}
