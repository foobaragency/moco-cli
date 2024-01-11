package cmd

import (
	"fmt"
	"log"
	"moco/data"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
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
	activitiesCmd.Flags().IntP("delete", "d", 0, "Delete activity by ID")
	rootCmd.AddCommand(activitiesCmd)
}
