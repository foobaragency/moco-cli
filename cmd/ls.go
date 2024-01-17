/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"moco/data"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List existing moco items (projects, tasks, activities)",
	Run: func(cmd *cobra.Command, args []string) {
        projects, _ := cmd.Flags().GetBool("projects")
        tasks, _ := cmd.Flags().GetBool("tasks")
        activities, _ := cmd.Flags().GetBool("activities")
        if (projects) {
            fmt.Println("Projects:")
            projects, err := data.GetProjects()
            if err != nil {
                fmt.Println(err)
            }
            for _, project := range projects {
                fmt.Printf("%d. %s\n", project.Id, project.Name)
            }
        } else if (tasks) {
            fmt.Println("Tasks:")
            projects, err := data.GetProjects()
            if err != nil {
                fmt.Println(err)
            }
            for _, project := range projects {
                tasks := project.Tasks
                for _, task := range tasks {
                    fmt.Printf("%d. %s\n", task.Id, task.Name)
                }
            }
        } else if (activities) {
            activities, err := data.GetActivities()
            if err != nil {
                fmt.Println(err)
            }
            for _, activity := range activities {
                fmt.Printf("%d. %s\n", activity.Id, activity.Description)
            }
        }
	},
}

func init() {
    lsCmd.Flags().BoolP("projects", "p", false, "List projects")
    lsCmd.Flags().BoolP("tasks", "t", false, "List tasks")
    lsCmd.Flags().BoolP("activities", "a", false, "List activities")
	rootCmd.AddCommand(lsCmd)
}
