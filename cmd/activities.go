package cmd

import (
	"fmt"
	"log"
	"moco/data"
	"moco/ui"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var activitiesCmd = &cobra.Command{
	Use:   "activities",
	Short: "List available activities",
	Run: func(cmd *cobra.Command, args []string) {
        if n, _ := cmd.Flags().GetBool("new"); n {
            projects, err := data.GetProjects()
            if err != nil {
                log.Fatal(err)
            }
            pickableProjects := make([]data.Pickable, len(projects))
            for i, project := range projects {
                pickableProjects[i] = project
            }
            selectedProject, err := ui.Pick(pickableProjects, "Select project")
            if err != nil {
                fmt.Println("Could not pick project")
                return
            }
            tasks := selectedProject.(data.Project).Tasks
            pickableTasks := make([]data.Pickable, len(tasks))
            for i, task := range tasks {
                pickableTasks[i] = task
            }
            selectedTask, err := ui.Pick(pickableTasks, "Select task")
            if err != nil {
                fmt.Println("Could not pick task")
                return
            }

            // fmt.Printf("Selected project: %s\n", selectedProject.(data.Project).Name)
            // fmt.Printf("Selected task: %s\n", selectedTask.(data.Task).Name)
            description, err := ui.Prompt("Enter activity description")
            err = data.CreateActivity(selectedProject.(data.Project).Id, selectedTask.(data.Task).Id, description)
            return
        }

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
            runningIndicator := " "
            if activity.TimerStartedAt != "" {
                runningIndicator = "*"
            }
            duration := time.Duration(activity.Seconds * 1000000000)
            activityNames = append(activityNames, fmt.Sprintf("%s%d\t%s\t%s", runningIndicator, activity.Id, duration.String(), activity.Description))
		}
		if len(activityNames) == 0 {
			fmt.Println("No activities found")
			return
		}
		fmt.Println(strings.Join(activityNames, "\n"))
	},
}

func init() {
    activitiesCmd.Flags().BoolP("new", "n", false, "Create a new activity")

	activitiesCmd.Flags().IntP("delete", "x", 0, "Delete activity by ID")
    activitiesCmd.Flags().IntP("edit", "e", 0, "Edit activity by ID")
    activitiesCmd.Flags().IntP("time", "t", 0, "Set the time for the activity (in seconds)")
    activitiesCmd.Flags().StringP("description", "d", "", "Set the description for the activity")

	rootCmd.AddCommand(activitiesCmd)
}
