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

var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Start tracking a new activity",
	Run: func(cmd *cobra.Command, args []string) {
        projects, err := data.GetProjects()
        if err != nil {
            fmt.Println("Could not retrieve projects", err)
            return
        }
            
        // flags get highest priority
		projectId, _ := cmd.Flags().GetInt("project")
		taskId, _ := cmd.Flags().GetInt("task")
		description, _ := cmd.Flags().GetString("description")

        // env has second priority
        err = godotenv.Load(".moco")
        if projectId == 0 && err == nil {
            projectId, err = strconv.Atoi(os.Getenv("MOCO_PROJECT_ID"))
            fmt.Println("read projectId", projectId)
        }
        if taskId == 0 && err == nil {
            taskId, err = strconv.Atoi(os.Getenv("MOCO_TASK_ID"))
            fmt.Println("read taskId", taskId)
        }

        // if no flags, prompt
		if projectId == 0 || err != nil {
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

func init() {
    trackCmd.Flags().IntP("project", "p", 0, "Set the project for the activity")
    trackCmd.Flags().IntP("task", "t", 0, "Set the task for the activity")
    trackCmd.Flags().StringP("description", "d", "", "Set the description for the activity")

	rootCmd.AddCommand(trackCmd)
}
