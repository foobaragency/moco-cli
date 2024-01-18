package cmd

import (
	"fmt"
	"moco/data"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new activity",
	Run: func(cmd *cobra.Command, args []string) {
		projectId, _ := cmd.Flags().GetInt("project")
		taskId, _ := cmd.Flags().GetInt("task")
		description, err := cmd.Flags().GetString("description")

		projects, err := data.GetProjects()
		if err != nil {
			fmt.Println("Could not retrieve projects", err)
		}

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
			huh.NewInput().Title("Description:").Prompt(">").Value(&description).Run()
		}

		err = data.CreateActivity(projectId, taskId, description)
		if err != nil {
			fmt.Println("Could not create activity:", err)
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
		seconds, err := cmd.Flags().GetInt("time")
		description, err := cmd.Flags().GetString("description")
		if err != nil || (seconds == 0 && description == "") {
			cmd.Help()
			return
		}
		err = data.EditActivity(activityId, seconds, description)
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
	editCmd.Flags().IntP("time", "t", 0, "Set the time for the activity (in seconds)")
	editCmd.Flags().StringP("description", "d", "", "Set the description for the activity")

	activitiesCmd.AddCommand(editCmd)

	newCmd.Flags().Bool("no-start", false, "Don't start the activity when created")
	activitiesCmd.AddCommand(newCmd)

	activitiesCmd.AddCommand(deleteCmd)

	rootCmd.AddCommand(activitiesCmd)
}
