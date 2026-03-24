package ui

import (
	"fmt"
	"log"

	"github.com/abhinav0411/spctl/models"
	"github.com/abhinav0411/spctl/spotify"
	tea "github.com/charmbracelet/bubbletea"
)

type spctl struct {
	currentScreen string
	client        *models.Client
	loginModel    *login
	screenModel   *screen
	logged_in     bool
	currentSong   models.CurrentSong
	device        []models.PlayerDevice
}

func initialModel() spctl {
	return spctl{
		currentScreen: "login",
		logged_in:     false,
		loginModel:    NewLogin(),
		screenModel: &screen{
			Player: newPlayer(),
		},
	}
}

func (m spctl) Init() tea.Cmd {
	return nil
}

func (m spctl) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}

	switch m.currentScreen {
	case "login":
		m.client, m.logged_in = m.loginModel.LoginUpdate(msg)
	case "screen":
		cmd = m.screenModel.ScreenUpdate(msg, m.currentSong, *m.client, m.device[0])
	}

	if m.logged_in {
		m.currentScreen = "screen"
		m.currentSong = spotify.GetCurrentSong(m.client)
		var err error
		m.device, err = spotify.GetDevice(m.client)
		if err != nil {
			log.Fatal("error while getting device")
		}
	}

	return m, cmd
}
func (m spctl) View() string {
	switch m.currentScreen {
	case "login":
		return m.loginModel.LoginView()
	case "screen":
		return m.screenModel.ScreenView()
	}
	return ""
}

func Start() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting the program")
	}
}
