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
		p.height = msg.Height - 5

	case tea.KeyMsg:
		if !p.focused {
			return p, nil
		}

		switch msg.String() {
		case "down", "j":
			if p.selected < len(p.items)-1 {
				p.selected++

				if p.selected >= p.scrollOffset+p.visibleItems() {
					p.scrollOffset++
				}
			}

		case "up", "k":
			if p.selected > 0 {
				p.selected--

				if p.selected < p.scrollOffset {
					p.scrollOffset--
				}
			}

		case "v":
			if len(p.items) > 0 {
				selected := p.items[p.selected]
				return p, func() tea.Msg {
					return SelectPlaylistMsg{
						ID:   selected.ID,
						URI:  selected.URI,
						Play: false,
					}
				}
			}

		case "p":
			if len(p.items) > 0 {
				selected := p.items[p.selected]
				return p, func() tea.Msg {
					return SelectPlaylistMsg{
						ID:   selected.ID,
						URI:  selected.URI,
						Play: true,
					}
				}
			}
		}
	}

	return p, nil
}

func (p Playlist) View() string {
	borderColor := lipgloss.Color("#1a1a1a")
	if p.focused {
		borderColor = lipgloss.Color("#1DB954")
	}

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#1DB954")).
		PaddingLeft(1)

	selectedNameStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#1DB954")).
		Bold(true)

	normalNameStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#cccccc"))

	selectedRowStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#0d2b18")).
		BorderLeft(true).
		BorderForeground(lipgloss.Color("#1DB954")).
		BorderStyle(lipgloss.ThickBorder()).
		PaddingLeft(1)

	normalRowStyle := lipgloss.NewStyle().
		PaddingLeft(2)

	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		Width(p.width - 2).
		Height(p.height - 2)

	if len(p.items) == 0 {
		empty := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#444444")).
			PaddingLeft(1).
			Render("no playlists found")
		return style.Render(empty)
	}

	start := p.scrollOffset
	end := start + p.visibleItems()
	if end > len(p.items) {
		end = len(p.items)
	}

	var content string
	content += titleStyle.Render("Playlists") + "\n\n"

	for i := start; i < end; i++ {
		item := p.items[i]
		if i == p.selected {
			line := "● " + selectedNameStyle.Render(item.Name)
			content += selectedRowStyle.Render(line) + "\n"
		} else {
			line := "○ " + normalNameStyle.Render(item.Name)
			content += normalRowStyle.Render(line) + "\n"
		}
	}

	return style.Render(content)
}

func (p Playlist) visibleItems() int {
	items := (p.height - 6)
	if items < 1 {
		return 1
	}
	return items
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
