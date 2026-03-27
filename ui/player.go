package ui

import (
	"strings"
	"time"

	"github.com/abhinav0411/spctl/models"
	"github.com/abhinav0411/spctl/spotify"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 42
)

var trackTitleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#e0aaff")).
	Bold(true).
	Render

type tickMsg time.Time

type player struct {
	isPlaying    bool
	percent      float64
	bar          progress.Model
	deltaDur     float64
	name_of_song string
	skip         string
	prev         string
	width        int
}

func NewPlayer() player {
	prog := progress.New(
		progress.WithWidth(44),
		progress.WithoutPercentage(),
		progress.WithScaledGradient("#20002c", "#9c27b0"),
	)
	return player{
		bar:  prog,
		skip: ">>",
		prev: "<<",
	}
}

func (m *player) PlayerUpdate(msg tea.Msg, current_song models.CurrentSong, client models.Client, device models.PlayerDevice) tea.Cmd {
	if current_song.Item.Name != "" {
		m.name_of_song = current_song.Item.Name
		m.isPlaying = current_song.IsPlaying
		m.deltaDur = 1000.0 / float64(current_song.Item.DurationMs)
		m.percent = float64(current_song.ProgressMs) / float64(current_song.Item.DurationMs)
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.bar.Width = msg.Width - padding*2 - 20
		if m.bar.Width > maxWidth {
			m.bar.Width = maxWidth
		}
		if m.bar.Width < 10 {
			m.bar.Width = 10
		}
		return nil

	case tickMsg:
		if m.isPlaying {
			m.percent += m.deltaDur
			if m.percent >= 1.0 {
				m.percent = 1.0
			}
		}
		return tickCmd()

	case progress.FrameMsg:
		updateModel, cmd := m.bar.Update(msg)
		m.bar = updateModel.(progress.Model)
		return cmd

	case tea.KeyMsg:
		if msg.String() == ">" {
			spotify.SkipNext(&client, device.ID)
		} else if msg.String() == "<" {
			spotify.SkipPrev(&client, device.ID)
		} else if msg.String() == " " && current_song.IsPlaying {
			spotify.Pause(&client, device.ID)
		} else if msg.String() == " " && !current_song.IsPlaying {
			spotify.TransferPlayback(&client, device.ID)
			spotify.Resume(&client, device.ID)
		}
	}
	return nil
}

func (m *player) PlayerView() string {
	pad := strings.Repeat(" ", padding)

	name := m.name_of_song

	maxTitleWidth := 20
	if len(name) > maxTitleWidth {
		name = name[:maxTitleWidth-3] + "..."
	}

	barWidth := m.bar.Width
	if barWidth == 0 {
		barWidth = maxWidth
	}

	controls := lipgloss.NewStyle().
		Width(barWidth + maxTitleWidth).
		Align(lipgloss.Center).
		Render(m.prev + "     " + m.skip)

	bar := m.bar.ViewAs(m.percent)
	title := trackTitleStyle(name)

	row := lipgloss.JoinHorizontal(
		lipgloss.Left,
		title,
		"  ",
		bar,
	)

	view := "\n"
	view += pad + row + "\n"
	view += pad + controls

	return view
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
