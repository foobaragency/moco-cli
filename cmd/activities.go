package cmd

import (
	"fmt"
	"moco/data"
	"strings"

	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var activitiesCmd = &cobra.Command{
	Use:   "activities",
	Short: "List activites assigned to you",
	Run: func(cmd *cobra.Command, args []string) {
        activites := data.GetActivities()
        var activityNames []string
        for _, activity := range activites {
            activityNames = append(activityNames, fmt.Sprintf("%d %s", activity.Id, activity.Description))
        }

        fmt.Println(strings.Join(activityNames, "\n"))
	},
}
