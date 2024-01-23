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
        if lines, err := cmd.Flags().GetBool("lines"); err == nil && lines {
            for _, project := range projects {
                fmt.Printf("%d %s\n", project.Id, project.GetName())
            }
            return
        }
        t := table.New().
            Border(lipgloss.RoundedBorder()).
            Headers("ID", "Name").
            StyleFunc(func(row, col int) lipgloss.Style {
                return lipgloss.NewStyle().Padding(0, 1)
            })
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
            Border(lipgloss.RoundedBorder()).
            StyleFunc(func(row, col int) lipgloss.Style {
                return lipgloss.NewStyle().Padding(0, 1)
            }).
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
        if lines, err := cmd.Flags().GetBool("lines"); err == nil && lines {
            for _, project := range projects {
                for _, task := range project.Tasks {
                    fmt.Printf("%d\t%s\n", task.Id, task.Name)
                }
            }
            return
        }
        for _, project := range projects {
            tasks := project.Tasks
            for _, task := range tasks {
                t.Row(fmt.Sprintf("%d", task.Id), task.Name)
            }
        }
        fmt.Println(t)
    },
}

func fmtDuration(d time.Duration) string {
    d = d.Round(time.Minute)
    h := d / time.Hour
    d -= h * time.Hour
    m := d / time.Minute
    return fmt.Sprintf("%02d:%02d", h, m)
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
                time := fmtDuration(duration)
                desc := activity.Description
                rows = append(rows, []string{id, date, time, desc})
            }
        }
        if lines, err := cmd.Flags().GetBool("lines"); err == nil && lines {
            for i := 0; i < len(rows); i++ {
                for j := 0; j < len(rows[i]); j++ {
                    fmt.Print(rows[i][j])
                    if j < len(rows[i])-1 {
                        fmt.Print("\t")
                    }
                }
                fmt.Println()
            }
            return
        }
        t := table.New().
            Border(lipgloss.RoundedBorder()).
            Headers("ID", "Date", "Time", "Description").
            StyleFunc(func(row int, col int) lipgloss.Style {
                style := lipgloss.NewStyle().Padding(0, 1)
                if (runningIndex < 1) {
                    return style
                }

                if runningIndex == row {
                    return style.Foreground(lipgloss.Color("9")).Bold(true)
                }
                return style
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
    projectCmd.Flags().BoolP("lines", "l", false, "List in lines")
    lsCmd.AddCommand(projectCmd)

    taskCmd.Flags().IntP("project", "p", 0, "Project ID")
    taskCmd.Flags().BoolP("lines", "l", false, "List in lines")
    lsCmd.AddCommand(taskCmd)

    activityCmd.Flags().BoolP("today", "t", false, "List tasks for today")
    activityCmd.Flags().BoolP("lines", "l", false, "List in lines")
    lsCmd.AddCommand(activityCmd)


	rootCmd.AddCommand(lsCmd)
}
