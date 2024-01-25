package cmd

import (
	"fmt"
	"moco/data"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new activity",
	Run: func(cmd *cobra.Command, args []string) {
        projects, err := data.GetProjects()
        if err != nil {
            fmt.Println("Could not retrieve projects", err)
            return
        }

        // flags has highest priority
		projectId, err := cmd.Flags().GetInt("project")
		taskId, err := cmd.Flags().GetInt("task")
        minutes, err := cmd.Flags().GetInt("minutes")
		description, err := cmd.Flags().GetString("description")

        // env has second priority
        err = godotenv.Load(".moco")
        if projectId == 0 && err == nil {
            projectId, err = strconv.Atoi(os.Getenv("MOCO_PROJECT_ID"))
        }
        if taskId == 0 && err == nil {
            taskId, err = strconv.Atoi(os.Getenv("MOCO_TASK_ID"))
        }

        // if no flags or config, prompt
		if projectId == 0 {
			options := make([]huh.Option[int], len(projects))
			for i, p := range projects {
				options[i] = huh.NewOption[int](p.Name, p.Id)
			}

			pform := huh.NewSelect[int]().Options(options...).Value(&projectId)
			pform.Run()
			if projectId == 0 {
				return
			}
		}

		var project data.Project
		for _, p := range projects {
			if p.Id == projectId {
				project = p
			}
		}

		if taskId == 0 {
			options := make([]huh.Option[int], len(project.Tasks))
			for i, t := range project.Tasks {
				options[i] = huh.NewOption[int](t.Name, t.Id)
			}
			tform := huh.NewSelect[int]().Options(options...).Value(&taskId)
			tform.Run()
			if taskId == 0 {
				return
			}
		}

        if description == "" {
            huh.NewInput().Title("Description:").Prompt("> ").Value(&description).Run()
        }

        if minutes == 0 {
            var minutesStr string
            huh.NewInput().Title("Time (minutes, entering '0' or empty will start a timer):").Prompt("> ").Value(&minutesStr).Run()
            if minutesStr == "" {
                minutes = 0
            } else {
                minutes, err = strconv.Atoi(minutesStr)
                if err != nil {
                    fmt.Println("Invalid minutes")
                    return
                }
            }
        }


		err = data.CreateActivity(projectId, taskId, description, minutes)
		if err != nil {
			fmt.Println("Could not create activity:", err)
		}
	},
}

var restartCmd = &cobra.Command{
    Use:   "restart",
    Short: "Restart an activity (empty <activity> will restart last activity)",
    Run: func(cmd *cobra.Command, args []string) {
        activityId, err := cmd.Flags().GetInt("activity")

        if activityId == 0 || err != nil {
            activities, err := data.GetActivities()
            if err != nil {
                fmt.Println("Could not retrieve activities", err)
                return
            }
            if len(activities) == 0 {
                fmt.Println("No activities found")
                return
            }
            activityId = activities[len(activities)-1].Id
        }

        err = data.RestartActivity(activityId)
        if err != nil {
            fmt.Println("Could not restart activity:", err)
        }
    },
}

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

    createCmd.Flags().IntP("project", "p", 0, "Set the project for the activity")
    createCmd.Flags().IntP("task", "t", 0, "Set the task for the activity")
    createCmd.Flags().StringP("description", "d", "", "Set the description for the activity")
    createCmd.Flags().IntP("minutes", "m", 0, "Set the number of minutes for the activity")

    restartCmd.Flags().IntP("activity", "a", 0, "Set the activity to restart")

	activitiesCmd.AddCommand(editCmd)
	activitiesCmd.AddCommand(createCmd)
	activitiesCmd.AddCommand(deleteCmd)
    activitiesCmd.AddCommand(restartCmd)

	rootCmd.AddCommand(activitiesCmd)
}
