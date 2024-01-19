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

        rows := make([][]string, 0)
        runningIndex := 0
        for _, activity := range activities {
            if !today || activity.Date == time.Now().Format("2006-01-02") {
                duration := time.Duration(activity.Seconds * 1000000000)
                if activity.TimerStartedAt != "" {
                    started, _ := time.Parse("2006-01-02T15:04:05Z", activity.TimerStartedAt)
                    elapsed := time.Since(started).Round(time.Second)
                    duration += elapsed
                    runningIndex = len(rows) + 1
                }
                id := fmt.Sprintf("%d", activity.Id)
                date := activity.Date
                time := duration.String()
                desc := activity.Description
                rows = append(rows, []string{id, date, time, desc})
            }
        }
        t := table.New().
            Border(lipgloss.NormalBorder()).
            Headers("ID", "Date", "Time", "Description").
            StyleFunc(func(row int, col int) lipgloss.Style {
                if (runningIndex < 1) {
                    return lipgloss.Style{}
                }

                if runningIndex == row {
                    return lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true)
                }
                return lipgloss.Style{}
            }).
            Rows(rows...)
        fmt.Println(t)
    },
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List existing moco items (projects, tasks, activities)",
}

func init() {
    lsCmd.AddCommand(projectCmd)

    taskCmd.Flags().IntP("project", "p", 0, "Project ID")
    lsCmd.AddCommand(taskCmd)

    activityCmd.Flags().BoolP("today", "t", false, "List tasks for today")
    lsCmd.AddCommand(activityCmd)

	rootCmd.AddCommand(lsCmd)
}
