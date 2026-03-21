package ui

import (
	"fmt"

	"github.com/abhinav0411/spctl/models"
	tea "github.com/charmbracelet/bubbletea"
)

type spctl struct {
	currentScreen string
	client        *models.Client
	loginModel    login
}

func initialModel() spctl {
	return spctl{
		currentScreen: "login",
	}
}

func (m spctl) Init() tea.Cmd {
	return nil
}

func (m spctl) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}

	switch m.currentScreen {
	case "login":
		m.loginModel.LoginUpdate(msg)
	}

	return m, nil
}
func (m spctl) View() string {
	switch m.currentScreen {
	case "login":
		return m.loginModel.LoginView()
	}
	return ""
}

func Start() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting the program")
	}
}
