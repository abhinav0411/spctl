package ui

import (
	"time"

	"github.com/abhinav0411/spctl/models"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type player struct {
	name_of_song string
	bar          progress.Model
	skip         string
	prev         string
}

type tickMsg time.Time

func (m *player) PlayerUpdate(msg tea.Msg, current_song models.CurrentSong) tea.Cmd {
	m.name_of_song = current_song.Item.Name
	time := (current_song.ProgressMs / current_song.Item.DurationMs) / 100

	switch msg := msg.(type) {
	case tickMsg:
		if m.bar.Percent() == 1.0 {
			return tea.Quit
		}

		cmd := m.bar.IncrPercent(float64(time))
		return cmd

	case progress.FrameMsg:
		var cmd tea.Cmd
		updateModel, cmd := m.bar.Update(msg)
		m.bar = updateModel.(progress.Model)
		return cmd

	default:
		return nil
	}
}

func (m *player) PlayerView() string {
	final_str := m.skip + " " + m.name_of_song + " " + m.prev + "\n" + m.bar.View()
	return final_str
}
