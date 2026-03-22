package ui

import (
	"github.com/abhinav0411/spctl/models"
	tea "github.com/charmbracelet/bubbletea"
)

type screen struct {
	Player      player
	spctl       string
	width       int
	initialized bool
}

func (m *screen) ScreenUpdate(msg tea.Msg, currentSong models.CurrentSong) tea.Cmd {
	if !m.initialized {
		m.initialized = true
		return tickCmd()
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
	}

	return m.Player.PlayerUpdate(msg, currentSong)
}

func (m *screen) ScreenView() string {
	m.spctl = `
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
	player_text := m.Player.PlayerView()

	final_str := m.spctl + player_text //lipgloss.NewStyle().Width(m.width).Align(lipgloss.Center).Render(m.spctl) + lipgloss.NewStyle().Render(player_text)
	return final_str
}
