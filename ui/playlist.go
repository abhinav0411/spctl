package ui

import (
	"github.com/abhinav0411/spctl/models"
	"github.com/abhinav0411/spctl/spotify"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type PlaylistItem struct {
	Name string
	ID   string
	URI  string
}

type SelectPlaylistMsg struct {
	ID   string
	URI  string
	Play bool
}

type Playlist struct {
	items        []PlaylistItem
	selected     int
	scrollOffset int
	width        int
	height       int
	focused      bool
}

func NewPlaylist() Playlist {
	return Playlist{
		items:    []PlaylistItem{},
		selected: 0,
		focused:  false,
	}
}

func (p Playlist) Update(msg tea.Msg) (Playlist, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.width = msg.Width / 3
		p.height = msg.Height - 8

	case tea.KeyMsg:
		if !p.focused {
			return p, nil
		}

		switch msg.String() {
		case "up", "k":
			if p.selected > 0 {
				p.selected--
				if p.selected < p.scrollOffset {
					p.scrollOffset = p.selected
				}
			}

		case "down", "j":
			if p.selected < len(p.items)-1 {
				p.selected++
				visibleLines := p.height - 4
				if p.selected >= p.scrollOffset+visibleLines {
					p.scrollOffset = p.selected - visibleLines + 1
				}
			}

		case "v":
			if len(p.items) > 0 {
				selected := p.items[p.selected]
				return p, func() tea.Msg {
					return SelectPlaylistMsg{
						ID:   selected.ID,
						URI:  selected.URI,
						Play: false, // view tracks
					}
				}
			}

		case "enter":
			if len(p.items) > 0 {
				selected := p.items[p.selected]
				return p, func() tea.Msg {
					return SelectPlaylistMsg{
						ID:   selected.ID,
						URI:  selected.URI,
						Play: true, // play directly
					}
				}
			}
		}
	}

	return p, nil
}

func (p Playlist) View() string {
	borderColor := lipgloss.Color("#555555")
	if p.focused {
		borderColor = lipgloss.Color("#1DB954") // Spotify green when focused
	}

	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		Width(p.width).
		Height(p.height)

	if len(p.items) == 0 {
		return style.Render("No playlists found")
	}

	visibleLines := p.height - 4
	if visibleLines < 1 {
		visibleLines = 1 // this saves it but shows only 1 line
	}

	end := p.scrollOffset + visibleLines
	if end > len(p.items) {
		end = len(p.items)
	}

	var content string
	for i := p.scrollOffset; i < end; i++ {
		item := p.items[i]
		line := item.Name
		if i == p.selected {
			line = "> " + line
		} else {
			line = "  " + line
		}
		content += line + "\n"
	}

	return style.Render(content)
}

func (p Playlist) SetPlaylists(data models.PlaylistResponse) Playlist {
	var items []PlaylistItem
	for _, pl := range data.Items {
		items = append(items, PlaylistItem{
			Name: pl.Name,
			ID:   pl.ID,
			URI:  pl.URI,
		})
	}
	p.items = items
	p.selected = 0
	p.scrollOffset = 0
	return p
}

func fetchPlaylistsCmd(client *models.Client) tea.Cmd {
	return func() tea.Msg {
		return spotify.GetUserPlaylists(client)
	}
}

func fetchPlaylistTracksCmd(client *models.Client, playlistID string) tea.Cmd {
	return func() tea.Msg {
		return spotify.GetPlaylistTracks(client, playlistID)
	}
}

func playPlaylistCmd(client models.Client, uri string, deviceID string) tea.Cmd {
	return func() tea.Msg {
		spotify.PlayPlaylist(&client, uri, deviceID)
		return nil
	}
}
