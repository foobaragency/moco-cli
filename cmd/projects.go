package cmd

import (
	"fmt"
	"moco/data"
	"strings"

	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
    Use:   "projects",
    Short: "List projects assigned to you",
    Run: func(cmd *cobra.Command, args []string) {
        projects := data.GetProjects()
        var projectNames []string
        for _, project := range projects {
            projectNames = append(projectNames, fmt.Sprintf("%d %s", project.Id, project.Name))
        }
        fmt.Println(strings.Join(projectNames, "\n"))
    },
}
