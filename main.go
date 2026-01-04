package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	username = iota
	password
)

const (
	glossPink = lipgloss.Color("#FF06B7")
	neon      = lipgloss.Color("#06c5ffff")
)

var (
	inputStyle   = lipgloss.NewStyle().Foreground(glossPink)
	headingStyle = lipgloss.NewStyle().Bold(true).Foreground(neon).Underline(true).Width(19).Align(lipgloss.Center)
)

type login struct {
	inputs  []textinput.Model
	focused int
	err     error
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func (m login) Init() tea.Cmd {
	return textinput.Blink
}

func initialModel() login {
	var inputs []textinput.Model = make([]textinput.Model, 2)
	inputs[username] = textinput.New()
	inputs[username].Width = 30
	inputs[username].Focus()

	inputs[password] = textinput.New()
	inputs[password].Width = 30

	return login{
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m login) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))
	fmt.Printf("%v", tea.WindowSize())

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				return m, tea.Quit
			}
			m.nextInput()

		case tea.KeyCtrlC:
			return m, tea.Quit

		case tea.KeyShiftTab:
			m.prevInput()

		case tea.KeyTab:
			m.nextInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

	}
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m login) View() string {
	return fmt.Sprintf(`%s
	%s %s

	%s %s
	
	`,
		headingStyle.Render("Welcome to spctl"),
		inputStyle.Width(0).Render("Username"),
		m.inputs[username].View(),
		inputStyle.Width(0).Render("Password"),
		m.inputs[password].View()) + "\n"
}

func (m *login) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

func (m *login) prevInput() {
	m.focused--

	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
