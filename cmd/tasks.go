package cmd

import (
	"fmt"
	"log"
	"moco/data"
	"strings"

	"github.com/spf13/cobra"
)

var tasksCmd = &cobra.Command{
    Use:   "tasks",
    Short: "List tasks for a given project",
    Run: func(cmd *cobra.Command, args []string) {
        projectId, _ := cmd.Flags().GetInt("project")
        project, err := data.GetProject(projectId)
        if err != nil {
            log.Fatal(err)
        }
        tasks := project.Tasks
        var taskNames []string
        for _, task := range tasks {
            taskNames = append(taskNames, fmt.Sprintf("%d %s", task.Id, task.Name))
        }
        if len(taskNames) == 0 {
            fmt.Println("No tasks found")
            return
        }
        fmt.Println(strings.Join(taskNames, "\n"))

    },
}

func init() {
    tasksCmd.Flags().IntP("project", "p", 0, "Project ID")
    tasksCmd.MarkFlagRequired("project")
    rootCmd.AddCommand(tasksCmd)
}
