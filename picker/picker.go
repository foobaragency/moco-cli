package picker

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string
	cursor   int
	header   string
	selected string
	exited   bool
}

func newModel(header string, choices []string) model {
	return model{
		choices:  choices,
		header:   header,
		cursor:   0,
		selected: "",
		exited:   false,
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
			m.exited = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
    s := m.header + "\n"
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

func Picker(header string, choices []string) (string, bool) {
    p, _ := tea.NewProgram(newModel(header, choices)).Run()
    return p.(model).selected, p.(model).exited
}
