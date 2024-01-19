package cmd

import (
	"fmt"
	"moco/data"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
    Use:   "project",
    Short: "List projects",
    Run: func(cmd *cobra.Command, args []string) {
        projects, _ := data.GetProjects()
        t := table.New().
            Border(lipgloss.NormalBorder()).
            Headers("ID", "Name")
        for _, project := range projects {
            t.Row(fmt.Sprintf("%d", project.Id), project.GetName())
        }
        fmt.Println(t)
    },
}

var taskCmd = &cobra.Command{
    Use:   "task",
    Short: "List tasks",
    Run: func(cmd *cobra.Command, args []string) {
        projectId, err := cmd.Flags().GetInt("project")

        t := table.New().
            Border(lipgloss.NormalBorder()).
            Headers("ID", "Name")
        if projectId != 0 && err == nil {
            project, err := data.GetProject(projectId)
            if err != nil {
                fmt.Println(err)
                return
            }
            for _, task := range project.Tasks {
                t.Row(fmt.Sprintf("%d", task.Id), task.Name)
            }
            fmt.Println(t)
            return
        }

        projects, _ := data.GetProjects()
        for _, project := range projects {
            tasks := project.Tasks
            for _, task := range tasks {
                t.Row(fmt.Sprintf("%d", task.Id), task.Name)
            }
        }
        fmt.Println(t)
    },
}

var activityCmd = &cobra.Command{
    Use:   "activity",
    Short: "List activities",
    Run: func(cmd *cobra.Command, args []string) {
        activities, _ := data.GetActivities()
        today, _ := cmd.Flags().GetBool("today")
        t := table.New().
            Border(lipgloss.NormalBorder()).
            Headers("ID", "Date", "Time", "Description")

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
                t.Row(fmt.Sprintf("%d", activity.Id), activity.Date, fmt.Sprintf("%s%s", prefix, duration.String()), activity.Description)
            }
        }
        fmt.Println(t)
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
