package ui

import (
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type sessionState int

const (
	listWidth  = 16
	listHeight = 16
)

const (
	search sessionState = iota
	playlist
	queue
)

type ScreenModel struct {
	state sessionState

	Searchbar textinput.Model
	Playlist  List
	Queue     List

	progress progress.Model

	logo string
}

func (m ScreenModel) Init() tea.Cmd {
	return nil
}

func NewScreen() ScreenModel {
	var playList, queueList List
	playList.initPlaylist(listWidth, listHeight)
	queueList.initQueue(listWidth, listHeight)

	return ScreenModel{
		state:    search,
		Queue:    queueList,
		Playlist: playList,
	}
}

func (m ScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m ScreenModel) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Left, m.Queue.View(), m.Playlist.View())
}
