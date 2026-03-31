package ui

import (
	"time"

	"github.com/abhinav0411/spctl/models"
	"github.com/abhinav0411/spctl/spotify"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg struct{}

type Player struct {
	width int

	progress progress.Model

	percent   float64
	isPlaying bool
	delta     float64

	trackName string
}

// --- INIT ---

func NewPlayer() Player {
	prog := progress.New(
		progress.WithDefaultGradient(),
		progress.WithoutPercentage(),
	)

	return Player{
		progress: prog,
		percent:  0,
		delta:    0.01,
	}
}

func (p Player) Init() tea.Cmd {
	return tickCmd()
}

// --- UPDATE ---

func (p Player) Update(msg tea.Msg, client *models.Client, id string) (Player, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		p.width = msg.Width

		barWidth := p.width / 3
		if barWidth < 20 {
			barWidth = 20
		}
		p.progress.Width = barWidth

	case tea.KeyMsg:
		switch msg.String() {

		case " ":
			if p.isPlaying {
				p.isPlaying = false
				return p, pauseCmd(client, id)
			} else {
				p.isPlaying = true
				return p, resumeCmd(client, id)
			}

		case "n":
			return p, nextCmd(client, id)

		case "p":
			return p, prevCmd(client, id)
		}

	case tickMsg:
		if p.isPlaying {
			p.percent += p.delta

			if p.percent >= 1.0 {
				p.percent = 1.0
			}
		}
		return p, tickCmd()
	}
	return p, nil
}

// --- SET SONG FROM BACKEND ---

func (p Player) SetSong(song models.CurrentSong) Player {
	p.trackName = song.Item.Name
	p.isPlaying = song.IsPlaying

	if song.Item.DurationMs > 0 {
		p.percent = float64(song.ProgressMs) / float64(song.Item.DurationMs)

		// increment per second
		p.delta = 1.0 / (float64(song.Item.DurationMs) / 1000.0)
	}

	return p
}

// --- VIEW ---

func (p Player) View() string {
	if p.width == 0 {
		return ""
	}

	left := lipgloss.NewStyle().
		Width(p.width / 3).
		Render(p.trackName)

	center := lipgloss.NewStyle().
		Width(p.width / 3).
		Render(p.progress.ViewAs(p.percent))

	right := lipgloss.NewStyle().
		Width(p.width / 3).
		Align(lipgloss.Right).
		Render("⏮    ⏭")

	bar := lipgloss.JoinHorizontal(
		lipgloss.Center,
		left,
		center,
		right,
	)

	container := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1)

	return container.Render(
		lipgloss.NewStyle().
			Width(p.width - container.GetHorizontalFrameSize()).
			Render(bar),
	)
}

// --- TICK LOOP ---

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

func PlayerTickCmd() tea.Cmd {
	return tickCmd()
}

func pauseCmd(client *models.Client, id string) tea.Cmd {
	return func() tea.Msg {
		spotify.Pause(client, id)
		return nil
	}
}

func resumeCmd(client *models.Client, id string) tea.Cmd {
	return func() tea.Msg {
		spotify.Resume(client, id)
		return nil
	}
}

func nextCmd(client *models.Client, id string) tea.Cmd {
	return func() tea.Msg {
		spotify.SkipNext(client, id)
		return nil
	}
}

func prevCmd(client *models.Client, id string) tea.Cmd {
	return func() tea.Msg {
		spotify.SkipPrev(client, id)
		return nil
	}
}
