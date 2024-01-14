package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func Prompt(prompt string) (string, error) {
	p := tea.NewProgram(initialModel(prompt))
	if r, err := p.Run(); err != nil {
        return "", err
	} else if m, ok := r.(PromptModel); ok {
        return m.textInput.Value(), nil
    } 
    return "", nil
}

type (
	errMsg error
)

type PromptModel struct {
	textInput textinput.Model
	err       error
}

func initialModel(prompt string) PromptModel {
	ti := textinput.New()
	ti.Placeholder = prompt
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return PromptModel{
		textInput: ti,
		err:       nil,
	}
}

func (m PromptModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m PromptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m PromptModel) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
