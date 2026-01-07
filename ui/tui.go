package ui

import tea "github.com/charmbracelet/bubbletea"

const (
	loginView sessionState = iota
	playlistView
	accountView
)

type tui struct {
	state  sessionState
	width  int
	height int
}

func (m tui) Init() tea.Cmd {
	return nil
}

func NewTuiModel() tui {
	return tui{
		state: loginView,
	}
}

// func (m tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	var cmd tea.Cmd

// }
