/*
Copyright © 2025
*/
package cmd

import (
	genai "chat-cli/clients"
	"chat-cli/data"
	"fmt"
	"log"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

type state int

const (
	stateSelectModel state = iota
	stateEnterPrompt
	stateLoading
	stateShowOutput
)

type item string

func (i item) Title() string       { return string(i) }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return string(i) }

type model struct {
	state       state
	list        list.Model
	textInput   textinput.Model
	spin        spinner.Model
	modelChoice string
	prompt      string
	response    string
	err         error
}

func initialModel() model {
	items := make([]list.Item, len(data.Models))
	for i, m := range data.Models {
		items[i] = item(m)
	}

	l := list.New(items, list.NewDefaultDelegate(), 30, 10)
	l.Title = "Choose a model"
	l.SetShowHelp(false)

	ti := textinput.New()
	ti.Focus()
	ti.Placeholder = "Enter a prompt..."

	sp := spinner.New()
	sp.Spinner = spinner.Dot

	return model{
		state:     stateSelectModel,
		list:      l,
		textInput: ti,
		spin:      sp,
	}
}

type resultMsg struct {
	output string
	err    error
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {

	// ----------------------
	// 1. Select Model
	// ----------------------
	case stateSelectModel:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				chosen := m.list.SelectedItem().(item)
				m.modelChoice = string(chosen)
				m.state = stateEnterPrompt
				return m, nil
			case "q", "esc":
				return m, tea.Quit
			}
		}
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd

	// ----------------------
	// 2. Enter Prompt
	// ----------------------
	case stateEnterPrompt:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				if m.textInput.Value() != "" {
					m.prompt = m.textInput.Value()
					m.state = stateLoading
					return m, tea.Batch(
						m.spin.Tick,
						func() tea.Msg {
							output, err := genai.Generate(m.modelChoice, m.prompt)
							return resultMsg{output: output, err: err}
						},
					)
				}
			case "esc":
				return m, tea.Quit
			}
		}
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd

	// ----------------------
	// 3. Loading Spinner
	// ----------------------
	case stateLoading:
		switch msg := msg.(type) {
		case resultMsg:
			m.response = msg.output
			m.err = msg.err
			m.state = stateShowOutput
			return m, nil
		}
		var cmd tea.Cmd
		m.spin, cmd = m.spin.Update(msg)
		return m, cmd

	// ----------------------
	// 4. Show Output
	// ----------------------
	case stateShowOutput:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "q" || msg.String() == "esc" {
				return m, tea.Quit
			}
		}
		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	switch m.state {

	case stateSelectModel:
		return m.list.View()

	case stateEnterPrompt:
		return fmt.Sprintf(
			"Model: %s\n\nEnter prompt:\n%s\n\n(press Enter to submit)",
			m.modelChoice,
			m.textInput.View(),
		)

	case stateLoading:
		return fmt.Sprintf(
			"%s Generating response...\n\n(press q to quit)",
			m.spin.View(),
		)

	case stateShowOutput:
		if m.err != nil {
			return fmt.Sprintf("Error: %v\n\nPress q to quit.", m.err)
		}
		rendered := markdown.Render(m.response, 80, 6)
		return fmt.Sprintf("%s\n\nPress q to quit.", string(rendered))
	}

	return ""
}

// ----------------------
// Cobra Command Wrapper
// ----------------------

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generate a response from a selected AI model (TUI)",
	Run: func(cmd *cobra.Command, args []string) {
		if err := tea.NewProgram(initialModel()).Start(); err != nil {
			log.Fatalf("Could not start TUI: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
