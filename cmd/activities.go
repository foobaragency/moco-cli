package cmd

import (
	"fmt"
	"moco/data"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit <activity>",
	Short: "Edit an activity",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		activityId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid activity id")
			return
		}
		minutes, err := cmd.Flags().GetInt("time")
		description, err := cmd.Flags().GetString("description")
		if err != nil || (minutes == 0 && description == "") {
			cmd.Help()
			return
		}
		err = data.EditActivity(activityId, minutes, description)
		if err != nil {
			fmt.Println("Could not edit activity:", err)
		}
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete <activity>",
	Short: "Delete an activity",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		activityId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid activity id")
			return
		}
		err = data.DeleteActivity(activityId)
		if err != nil {
			fmt.Println("Could not delete activity:", err)
		}
	},
}

var activitiesCmd = &cobra.Command{
	Use:   "activity",
	Short: "Create, edit, or delete activities",
}

func init() {
	editCmd.Flags().IntP("time", "t", 0, "Set the time for the activity (in minutes)")
	editCmd.Flags().StringP("description", "d", "", "Set the description for the activity")

	activitiesCmd.AddCommand(editCmd)
	activitiesCmd.AddCommand(deleteCmd)

	rootCmd.AddCommand(activitiesCmd)
}
