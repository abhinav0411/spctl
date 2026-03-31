package ui

import (
	"fmt"
	"time"

	"github.com/abhinav0411/spctl/models"
	"github.com/abhinav0411/spctl/spotify"
	tea "github.com/charmbracelet/bubbletea"
)

type spctl struct {
	currentScreen string

	client      *models.Client
	loginModel  *login
	screenModel *screen

	loggedIn    bool
	currentSong models.CurrentSong
	device      models.PlayerDevice
}

// --- COMMANDS ---

func refreshSongCmd(client *models.Client) tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return spotify.GetCurrentSong(client)
	})
}

func fetchDeviceCmd(client *models.Client) models.PlayerDevice {
	devices, err := spotify.GetDevice(client)
	if err != nil {
		return models.PlayerDevice{}
	}
	return devices[0]
}

// --- INITIAL MODEL ---

func initialModel() spctl {
	return spctl{
		currentScreen: "login",
		loggedIn:      false,
		loginModel:    NewLogin(),
		screenModel:   NewScreen(),
	}
}

// --- INIT ---

func (m spctl) Init() tea.Cmd {
	return nil
}

// --- UPDATE ---

func (m spctl) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}

	case models.CurrentSong:
		m.currentSong = msg

		if m.client != nil {
			return m, refreshSongCmd(m.client)
		}
	case []models.PlayerDevice:
		if len(msg) > 0 {
			m.device = msg[0]
		}
	}

	switch m.currentScreen {

	case "login":
		m.client, m.loggedIn = m.loginModel.LoginUpdate(msg)

		if m.loggedIn {
			m.currentScreen = "screen"

			if m.client != nil {
				m.device = fetchDeviceCmd(m.client)
				return m, tea.Batch(
					refreshSongCmd(m.client),
					PlayerTickCmd(),
				)
			}
		}

	case "screen":
		if m.client != nil {
			cmd = m.screenModel.ScreenUpdate(
				msg,
				m.currentSong,
				*m.client,
				m.device,
			)
		}
	}

	return m, cmd
}

// --- VIEW ---

func (m spctl) View() string {
	switch m.currentScreen {
	case "login":
		return m.loginModel.LoginView()
	case "screen":
		return m.screenModel.ScreenView()
	}
	return ""
}

// --- START ---

func Start() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting the program")
	}
}
