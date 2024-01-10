package cmd

import (
	"fmt"
	"log"
	"moco/data"
	"strings"

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
        var activityNames []string
        for _, activity := range activites {
            activityNames = append(activityNames, fmt.Sprintf("%d %s", activity.Id, activity.Description))
        }
        if len(activityNames) == 0 {
            fmt.Println("No activities found")
            return
        }
        fmt.Println(strings.Join(activityNames, "\n"))
	},
}

func init() {
    rootCmd.AddCommand(activitiesCmd)
}
