package ui

import (
	"strings"

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

	styledLogo := style.Render(m.spctl)
	logo := lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		styledLogo,
	)

	player := lipgloss.NewStyle().Width(m.width).Align(lipgloss.Right).Render(player_text)

	spacerHeight := m.height - lipgloss.Height(logo) - lipgloss.Height(player)
	if spacerHeight < 0 {
		spacerHeight = 0
	}
	spacer := strings.Repeat("\n", spacerHeight)

	return logo + spacer + player
}
