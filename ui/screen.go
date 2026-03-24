package ui

import (
	"github.com/abhinav0411/spctl/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type screen struct {
	Player      player
	spctl       string
	height      int
	width       int
	initialized bool
}

func NewScreen() *screen {
	return &screen{
		Player: NewPlayer(),
		spctl: `
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
`,
	}
}

func (m *screen) ScreenUpdate(msg tea.Msg, currentSong models.CurrentSong, client models.Client, device models.PlayerDevice) tea.Cmd {
	if !m.initialized {
		m.initialized = true
		return tickCmd()
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m.Player.PlayerUpdate(msg, currentSong, client, device)
}

func (m *screen) ScreenView() string {
	player_text := m.Player.PlayerView()

	final_str := style.Width(m.width).Align(lipgloss.Center).Render(m.spctl) + lipgloss.NewStyle().Height(m.height).Width(m.width).AlignHorizontal(lipgloss.Bottom).AlignVertical(lipgloss.Right).Render(player_text)
	return final_str
}
