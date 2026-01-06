package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	username = iota
	password
)

const (
	glossPink = lipgloss.Color("201")
	neon      = lipgloss.Color("86")
)

var (
	inputStyle   = lipgloss.NewStyle().Foreground(glossPink)
	headingStyle = lipgloss.NewStyle().Bold(true).Foreground(neon)
)

type LoginModel struct {
	inputs []textinput.Model
	focus  int
	width  int
	err    error
}

func (m LoginModel) Init() tea.Cmd {
	return textinput.Blink
}

func NewLoginModel() LoginModel {
	var inputs []textinput.Model = make([]textinput.Model, 2)
	inputs[username] = textinput.New()
	inputs[username].Width = 30
	inputs[username].Focus()

	inputs[password] = textinput.New()
	inputs[password].Width = 30

	return LoginModel{
		inputs: inputs,
		focus:  0,
		err:    nil,
	}
}

func (m LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			if m.focus == len(m.inputs)-1 {
				return m, tea.Quit
			}
			m.nextInput()

		case "q", "ctrl+c":
			return m, tea.Quit

		case "tab":
			m.nextInput()

		case "shift+tab":
			m.prevInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focus].Focus()

	}
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m LoginModel) View() string {
	return fmt.Sprintf(`%s
			
			%s %s
			
			%s %s
			
	`,
		headingStyle.Width(m.width).Align(lipgloss.Center).Render(`

                _   _
 ___ _ __   ___| |_| |
/ __| '_ \ / __| __| |
\__ \ |_) | (__| |_| |
|___/ .__/ \___|\__|_|
    |_|

`),
		inputStyle.Render("Username"),
		m.inputs[username].View(),
		inputStyle.Render("Password"),
		m.inputs[password].View()) + "\n"
}

func (m *LoginModel) nextInput() {
	m.focus = (m.focus + 1) % len(m.inputs)
}

func (m *LoginModel) prevInput() {
	m.focus--

	if m.focus < 0 {
		m.focus = len(m.inputs) - 1
	}
}
