package cmd

import (
	"fmt"
	"log"
	"moco/data"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var activitiesCmd = &cobra.Command{
	Use:   "activities",
	Short: "List available activities",
	Run: func(cmd *cobra.Command, args []string) {
		activites, err := data.GetActivities()
		if err != nil {
			log.Fatal(err)
		}

		if id, _ := cmd.Flags().GetInt("delete"); id != 0 {
			err := data.DeleteActivity(id)
			if err != nil {
				log.Fatal(err)
			}
			return
		}

        if edit, _ := cmd.Flags().GetInt("edit"); edit != 0 {
            seconds, err := cmd.Flags().GetInt("time")
            description, err := cmd.Flags().GetString("description")
            if err != nil || (seconds == 0 && description == "") {
                cmd.Help()
                return
            }
            err = data.EditActivity(edit, seconds, description)
            if err != nil {
                fmt.Println("Could not edit activity:", err)
            }
            return
            
        }

		var activityNames []string
		for _, activity := range activites {
			elapsedString := ""
			if activity.TimerStartedAt != "" {
				startedAtTime, err := time.Parse("2006-01-02T15:04:05Z", activity.TimerStartedAt)
				if err != nil {
					log.Fatal(err)
				}
				now := time.Now()
				elapsed := now.Sub(startedAtTime).Round(time.Second)
				elapsedString = fmt.Sprintf("(%s)", elapsed)
			}
			activityNames = append(activityNames, fmt.Sprintf("%d %s %s", activity.Id, activity.Description, elapsedString))
		}
		if len(activityNames) == 0 {
			fmt.Println("No activities found")
			return
		}
		fmt.Println(strings.Join(activityNames, "\n"))
	},
}

func init() {
	activitiesCmd.Flags().IntP("delete", "x", 0, "Delete activity by ID")
    activitiesCmd.Flags().IntP("edit", "e", 0, "Edit activity by ID")
    activitiesCmd.Flags().IntP("time", "t", 0, "Set the time for the activity (in seconds)")
    activitiesCmd.Flags().StringP("description", "d", "", "Set the time for the activity")

	rootCmd.AddCommand(activitiesCmd)
}
