/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"moco/data"
	"time"

	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
    Use:   "project",
    Short: "List projects",
    Run: func(cmd *cobra.Command, args []string) {
        projects, _ := data.GetProjects()
        for _, project := range projects {
            fmt.Printf("%d %s\n", project.Id, project.Name)
        }
    },
}

var taskCmd = &cobra.Command{
    Use:   "task",
    Short: "List tasks",
    Run: func(cmd *cobra.Command, args []string) {
        projectId, err := cmd.Flags().GetInt("project")

        if projectId != 0 && err == nil {
            project, err := data.GetProject(projectId)
            if err != nil {
                fmt.Println(err)
                return
            }
            for _, task := range project.Tasks {
                fmt.Printf("%d %s\n", task.Id, task.Name)
            }
            return
        }

        projects, _ := data.GetProjects()
        for _, project := range projects {
            tasks := project.Tasks
            for _, task := range tasks {
                fmt.Printf("%d %s\n", task.Id, task.Name)
            }
        }
    },
}

var activityCmd = &cobra.Command{
    Use:   "activity",
    Short: "List activities",
    Run: func(cmd *cobra.Command, args []string) {
        activities, _ := data.GetActivities()
        today, _ := cmd.Flags().GetBool("today")
        for _, activity := range activities {
            if !today || activity.Date == time.Now().Format("2006-01-02") {
                prefix := " "
                duration := time.Duration(activity.Seconds * 1000000000)
                if activity.TimerStartedAt != "" {
                    prefix = "*"
                    started, _ := time.Parse("2006-01-02T15:04:05Z", activity.TimerStartedAt)
                    elapsed := time.Since(started).Round(time.Second)
                    duration += elapsed
                }
                fmt.Printf("%s%d\t%s\t(%s)\n", prefix, activity.Id, activity.Description, duration.String())
            }
        }
    },
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List existing moco items (projects, tasks, activities)",
}

func init() {
    lsCmd.Flags().BoolP("projects", "p", false, "List projects")
    lsCmd.Flags().BoolP("tasks", "t", false, "List tasks")
    lsCmd.Flags().BoolP("activities", "a", false, "List activities")

    lsCmd.AddCommand(projectCmd)

    taskCmd.Flags().IntP("project", "p", 0, "Project ID")
    lsCmd.AddCommand(taskCmd)

    activityCmd.Flags().BoolP("today", "t", false, "List tasks for today")
    lsCmd.AddCommand(activityCmd)

	rootCmd.AddCommand(lsCmd)
}
