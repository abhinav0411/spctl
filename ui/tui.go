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
	if err != nil || len(devices) == 0 {
		return models.PlayerDevice{}
	}
	return devices[0]
}

func fetchDevicesCmd(client *models.Client) tea.Cmd {
	return func() tea.Msg {
		devices, err := spotify.GetDevice(client)
		if err != nil || len(devices) == 0 {
			return []models.PlayerDevice{}
		}
		return devices
	}
}

func fetchSongCmd(client *models.Client) tea.Cmd {
	return func() tea.Msg {
		return spotify.GetCurrentSong(client)
	}
}

// 🔥 slow tick
type slowTickMsg struct{}

func slowTickCmd() tea.Cmd {
	return tea.Tick(time.Second*2, func(t time.Time) tea.Msg {
		return slowTickMsg{}
	})
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
	return slowTickCmd()
}

// --- UPDATE ---

func (m spctl) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}

	case models.CurrentSong:
		m.currentSong = msg
		return m, nil // ← add this, don't pass song msg down to screen

	case []models.PlayerDevice:
		if len(msg) > 0 {
			m.device = msg[0]
		}
		return m, nil // ← same here

	case slowTickMsg:
		if m.loggedIn && m.client != nil {
			cmds = append(cmds,
				fetchSongCmd(m.client),
				fetchDevicesCmd(m.client),
			)
		}
		cmds = append(cmds, slowTickCmd())
		return m, tea.Batch(cmds...) // ← return early, don't pass tick down
	}

	// only keypresses and window/other msgs reach the screen
	switch m.currentScreen {
	case "login":
		m.client, m.loggedIn = m.loginModel.LoginUpdate(msg)
		if m.loggedIn {
			m.currentScreen = "screen"
			if m.client != nil {
				m.device = fetchDeviceCmd(m.client)
				cmds = append(cmds,
					PlayerTickCmd(),
					slowTickCmd(),
				)
			}
		}
	case "screen":
		if m.client != nil {
			cmd := m.screenModel.ScreenUpdate(
				msg,
				m.currentSong,
				*m.client,
				m.device,
			)
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
	}

	return m, tea.Batch(cmds...)
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
