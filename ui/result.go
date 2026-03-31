package ui

import (
	"github.com/abhinav0411/spctl/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TrackItem struct {
	Title  string
	Artist string
	URI    string
}

type PlaySongMsg struct {
	URI string
}

type Playlist struct {
	items    []TrackItem
	selected int
	width    int
	height   int
	focused  bool
}

func NewPlaylist() Playlist {
	return Playlist{
		items:    []TrackItem{},
		focused:  false,
		selected: 0,
	}
}

func (p Playlist) Update(msg tea.Msg) (Playlist, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		p.width = msg.Width / 3
		p.height = msg.Height - 8

	case tea.KeyMsg:
		switch msg.String() {

		case "up", "k":
			if p.selected > 0 {
				p.selected--
			}

		case "down", "j":
			if p.selected < len(p.items)-1 {
				p.selected++
			}

		case "enter":
			if len(p.items) > 0 {
				return p, func() tea.Msg {
					return PlaySongMsg{
						URI: p.items[p.selected].URI,
					}
				}
			}
		}
	}

	return p, nil
}

func (p Playlist) View() string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(p.width).
		Height(p.height)

	if len(p.items) == 0 {
		return style.Render("Search for a song...")
	}

	var lines []string

	for i, item := range p.items {
		line := item.Title + " - " + item.Artist

		if i == p.selected {
			line = "> " + line
		} else {
			line = "  " + line
		}

		lines = append(lines, line)
	}

	// 🔥 clamp to height
	if len(lines) > p.height-2 {
		lines = lines[:p.height-2]
	}

	content := ""
	for _, l := range lines {
		content += l + "\n"
	}

	return style.Render(content)
}
func (p Playlist) SetResults(search models.SearchResult) Playlist {
	var items []TrackItem

	for _, track := range search.Tracks.Items {
		artist := ""
		if len(track.Artists) > 0 {
			artist = track.Artists[0].Name
		}

		items = append(items, TrackItem{
			Title:  track.Name,
			Artist: artist,
			URI:    track.URI,
		})
	}

	p.items = items
	p.selected = 0

	return p
}
