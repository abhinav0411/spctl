package ui

import (
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
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
	Playlist  ListSong
	Queue     ListSong

	progress progress.Model

	logo string
}

func (m ScreenModel) Init() tea.Cmd {
	return nil
}

func NewScreen() ScreenModel {
	var songList List
	songList.initPlaylist(listWidth, listHeight)

	return ScreenModel{
		state: search,
		Queue: songList,
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
	return m.Queue.View()
}
