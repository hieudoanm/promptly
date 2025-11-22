/*
Copyright © 2025
*/
package cmd

import (
	"chat-cli/clients/chat"
	"chat-cli/data"
	"fmt"
	"log"
	"strings"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

type state int

const (
	stateSelectModel state = iota
	stateChat
	stateLoading
)

type item string

func (i item) Title() string       { return string(i) }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return string(i) }

type resultMsg struct {
	output string
	err    error
}

type model struct {
	state state

	list        list.Model
	modelChoice string

	history  viewport.Model
	messages []string

	input textarea.Model

	spin spinner.Model

	err error
}

// ---------------------------
// Lipgloss Styles
// ---------------------------

var (
	historyStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 1).
			Width(82)

	inputStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(0, 1).
			Width(82)
)

func initialModel() model {
	items := make([]list.Item, len(data.Models))
	for i, m := range data.Models {
		items[i] = item(m)
	}

	l := list.New(items, list.NewDefaultDelegate(), 30, 10)
	l.Title = "Choose a model"
	l.SetShowHelp(false)

	ti := textarea.New()
	ti.Placeholder = "Type a message..."
	ti.Focus()
	ti.CharLimit = 5000
	ti.SetHeight(3)

	vp := viewport.New(80, 20)
	vp.SetContent("")

	sp := spinner.New()
	sp.Spinner = spinner.Dot

	return model{
		state:   stateSelectModel,
		list:    l,
		input:   ti,
		history: vp,
		spin:    sp,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) appendMessage(role, msg string) {
	rendered := markdown.Render(msg, 80, 6)
	entry := fmt.Sprintf("[%s]\n%s\n", strings.ToUpper(role), string(rendered))

	m.messages = append(m.messages, entry)
	m.history.SetContent(strings.Join(m.messages, "\n"))
	m.history.GotoBottom()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {

	case stateSelectModel:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				chosen := m.list.SelectedItem().(item)
				m.modelChoice = string(chosen)
				m.state = stateChat
				return m, nil
			case "q", "esc", "ctrl+c":
				return m, tea.Quit
			}
		}
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd

	case stateChat:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {

			case "enter":
				text := strings.TrimSpace(m.input.Value())
				if text == "" {
					return m, nil
				}

				m.appendMessage("user", text)
				m.input.Reset()

				m.state = stateLoading
				return m,
					tea.Batch(
						m.spin.Tick,
						func() tea.Msg {
							output, err := chat.Generate(
								m.modelChoice,
								text,
							)
							return resultMsg{output: output, err: err}
						},
					)

			case "esc", "ctrl+c":
				return m, tea.Quit
			}
		}

		var cmds []tea.Cmd

		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		cmds = append(cmds, cmd)

		m.history, cmd = m.history.Update(msg)
		cmds = append(cmds, cmd)

		return m, tea.Batch(cmds...)

	case stateLoading:
		switch msg := msg.(type) {
		case resultMsg:
			if msg.err != nil {
				m.appendMessage("error", msg.err.Error())
			} else {
				m.appendMessage("assistant", msg.output)
			}

			m.state = stateChat
			return m, nil
		}

		var cmd tea.Cmd
		m.spin, cmd = m.spin.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	switch m.state {

	case stateSelectModel:
		return m.list.View()

	case stateLoading:
		return fmt.Sprintf(
			"%s Generating response...\n\nPress esc to quit.",
			m.spin.View(),
		)

	case stateChat:
		return fmt.Sprintf(
			"Model: %s\n\n%s\n\n%s\n\n(Enter to send, Esc to quit)",
			m.modelChoice,
			historyStyle.Render(m.history.View()),
			inputStyle.Render(m.input.View()),
		)
	}

	return ""
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Chat-style TUI with selected AI model",
	Run: func(cmd *cobra.Command, args []string) {
		if err := tea.NewProgram(initialModel()).Start(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
