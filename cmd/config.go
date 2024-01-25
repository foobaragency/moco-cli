package cmd

import (
	"fmt"
	"moco/data"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Add default options when running in current directory",
	Run: func(cmd *cobra.Command, args []string) {
        pid, perr := cmd.Flags().GetInt("project")
        tid, terr := cmd.Flags().GetInt("task")
        desc, derr := cmd.Flags().GetString("description")

        projects, err := data.GetProjects()
        if err != nil {
            fmt.Println("Could not retrieve projects", err)
            return
        }

		if pid == 0 || perr != nil {
			options := make([]huh.Option[int], len(projects))
			for i, p := range projects {
				options[i] = huh.NewOption[int](p.Name, p.Id)
			}

			pform := huh.NewSelect[int]().Options(options...).Value(&pid)
			pform.Run()
			if pid == 0 {
				return
			}
		}

		var project data.Project
		for _, p := range projects {
			if p.Id == pid {
				project = p
			}
		}

		if tid == 0 || terr != nil {
			options := make([]huh.Option[int], len(project.Tasks))
			for i, t := range project.Tasks {
				options[i] = huh.NewOption[int](t.Name, t.Id)
			}
			tform := huh.NewSelect[int]().Options(options...).Value(&tid)
			tform.Run()
			if tid == 0 {
				return
			}
		}

        if desc == "" || derr != nil {
            huh.NewInput().Title("Description:").Prompt("> ").Value(&desc).Run()
        }

        os.WriteFile(".moco", []byte(fmt.Sprintf("MOCO_PROJECT_ID=%d\nMOCO_TASK_ID=%d\nMOCO_DESCRIPTION=%s", pid, tid, desc)), 0644)
	},
}

func init() {
    configCmd.Flags().IntP("project", "p", 0, "Set the default project ID")
    configCmd.Flags().IntP("task", "t", 0, "Set the default task ID")
    configCmd.Flags().StringP("description", "d", "", "Set the default description")
	rootCmd.AddCommand(configCmd)
}
