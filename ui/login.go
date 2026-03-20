package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var green = lipgloss.Color("#1DB954")
var style = lipgloss.NewStyle().Foreground(green)
var login_style = lipgloss.NewStyle().Blink(true)

type login struct {
	width      int
	height     int
	login_Text string
}

func (m login) Init() tea.Cmd {
	return nil
}

func (m login) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m login) View() string {
	spctl := `
                             █████    ████ 
                            ░░███    ░░███ 
  █████  ████████   ██████  ███████   ░███ 
 ███░░  ░░███░░███ ███░░███░░░███░    ░███ 
░░█████  ░███ ░███░███ ░░░   ░███     ░███ 
 ░░░░███ ░███ ░███░███  ███  ░███ ███ ░███ 
 ██████  ░███████ ░░██████   ░░█████  █████
░░░░░░   ░███░░░   ░░░░░░     ░░░░░  ░░░░░ 
         ░███                              
         █████                             
        ░░░░░                              
`

	m.login_Text = `
 █   █▀█ █▀▀ ▀█▀ █▀█
 █   █ █ █ █  █  █ █
 ▀▀▀ ▀▀▀ ▀▀▀ ▀▀▀ ▀ ▀
 Press ENTER to login.
`

	final_str := style.Width(m.width).Align(lipgloss.Center).Render(spctl) + login_style.Width(m.width).PaddingTop((m.height-3)/2).Align(lipgloss.Center).Render(m.login_Text)
	return final_str
}

func Start() {
	p := tea.NewProgram(model{}, tea.WithAltScreen())
	p.Run()
}
