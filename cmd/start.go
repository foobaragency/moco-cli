package cmd

import (
	"fmt"
	"log"
    "moco/data"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var projectId int
var taskId int

type model struct {
	choices  []string
	cursor   int
	header   string
	selected string
}

func newModel(header string, choices []string) model {
	return model{
		choices:  choices,
		cursor:   0,
		selected: choices[0],
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "enter":
            m.selected = m.choices[m.cursor]
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
        case "q", "ctrl+c":
            return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
    s := "\n"
	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = ">" 
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\nq or ctrl+c to quit"

	return s
}

// rootCmd represents the base command when called without any subcommands
var startCmd = &cobra.Command{
	Use:   "start [id]",
	Short: "Start time tracking for a given project",
	Run: func(cmd *cobra.Command, args []string) {
        projects := data.GetProjects()
        
		if projectId == 0 {
			var projectNames []string
			for _, project := range projects {
				projectNames = append(projectNames, project.Name)
			}

			p, err := tea.NewProgram(newModel("Projects", projectNames)).Run()
			if err != nil {
				log.Fatal(err)
			}

			if p, ok := p.(model); ok {
				for _, project := range projects {
					if project.Name == p.selected {
						projectId = project.Id
					}
				}
			} else {
				log.Fatal("Failed to parse model")
			}
		} 
        var project data.Project
        projectFound := false
        for _, project = range projects {
            if project.Id == projectId {
                projectFound = true
                break
            }
        }
        if !projectFound {
            log.Fatal("Project not found")
        }

		if taskId == 0 {
            tasks := project.Tasks
            taskNames := make([]string, len(tasks))
            for i, task := range tasks {
                taskNames[i] = task.Name
            }


			p, err := tea.NewProgram(newModel("Tasks", taskNames)).Run()

			if err != nil {
				log.Fatal(err)
			}

			if p, ok := p.(model); ok {
				for _, task := range tasks {
					if task.Name == p.selected {
						taskId = task.Id
					}
				}
			} else {
				log.Fatal("Failed to parse model")
			}
		} 
        var task data.Task
        taskFound := false
        for _, task = range project.Tasks {
            if task.Id == taskId {
                taskFound = true
                break
            }
        }
        if !taskFound {
            log.Fatal("Task not found")
        }
        fmt.Printf("taskId: %v, projectId: %v", taskId, projectId)

        // make a patch request

	},
}

func init() {
	startCmd.Flags().IntVarP(&projectId, "project", "p", 0, "Project ID")
	startCmd.Flags().IntVarP(&taskId, "task", "t", 0, "Task ID")
	rootCmd.AddCommand(startCmd)
}
