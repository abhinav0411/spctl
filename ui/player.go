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

	trackName  string
	artistName string
}

func NewPlayer() Player {
	prog := progress.New(
		progress.WithGradient("#0d9488", "#0a0a0a"),
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

		case "b":
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

func (p Player) SetSong(song models.CurrentSong) Player {
	if song.Item.Name == "" {
		return p
	}
	p.trackName = song.Item.Name
	p.isPlaying = song.IsPlaying
	if len(song.Item.Artists) > 0 {
		p.artistName = song.Item.Artists[0].Name
	} else {
		p.artistName = "unknown artist"
	}
	if song.Item.DurationMs > 0 {
		p.percent = float64(song.ProgressMs) / float64(song.Item.DurationMs)
		p.delta = 1.0 / (float64(song.Item.DurationMs) / 1000.0)
	}
	return p
}

func (p Player) View() string {
	if p.width == 0 {
		return ""
	}

	trackStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#4fcac0")).
		Bold(true).
		MaxWidth(p.width / 4)

	artistStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#a24ead")).
		Italic(true)

	track := trackStyle.Render(p.trackName)
	artist := artistStyle.Render(p.artistName)

	trackBlock := track + "\n" + artist

	left := lipgloss.NewStyle().
		Width(p.width / 3).
		Render(trackBlock)

	center := lipgloss.NewStyle().
		Width(p.width / 3).
		Render(p.progress.ViewAs(p.percent))

	bar := lipgloss.JoinHorizontal(
		lipgloss.Center,
		left,
		center,
	)

	borderColor := lipgloss.Color("#444444")

	container := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		Background(lipgloss.Color("#111111")).
		Padding(0, 1)

	return container.Render(
		lipgloss.NewStyle().
			Width(p.width - container.GetHorizontalFrameSize()).
			Render(bar),
	)
}

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

func playSongCmd(client models.Client, uri string, deviceID string) tea.Cmd {
	return func() tea.Msg {

		body := map[string]interface{}{
			"uris": []string{uri},
		}

		spotify.Start(&client, body, deviceID)

		return nil
	}
}
