package ui

import (
	"github.com/abhinav0411/spctl/models"
	"github.com/abhinav0411/spctl/spotify"
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
	spctl      string
	returnMsg  tea.Msg
}

func NewLogin() *login {
	return &login{
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
		login_Text: `
 █   █▀█ █▀▀ ▀█▀ █▀█
 █   █ █ █ █  █  █ █
 ▀▀▀ ▀▀▀ ▀▀▀ ▀▀▀ ▀ ▀
 Press ENTER to login.
`,
	}
}

func (m *login) LoginUpdate(msg tea.Msg) (*models.Client, bool) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			msg, client := spotify.LoginCmd()
			m.returnMsg = msg
			return &client, true
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return &models.Client{}, false
}

func (m *login) LoginView() string {

	final_str := style.Width(m.width).Align(lipgloss.Center).Render(m.spctl) + login_style.Width(m.width).PaddingTop((m.height-3)/2).Align(lipgloss.Center).Render(m.login_Text)
	return final_str
}
