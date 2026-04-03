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

type Result struct {
	items        []TrackItem
	selected     int
	scrollOffset int
	width        int
	height       int
	focused      bool
}

func NewResult() Result {
	return Result{
		items:    []TrackItem{},
		selected: 0,
	}
}

func (p Result) Update(msg tea.Msg) (Result, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.width = msg.Width / 3
		p.height = msg.Height - 5 // 3 search + 5 player approx

	case tea.KeyMsg:
		// Only handle navigation keys if this panel is focused
		if !p.focused {
			return p, nil
		}

		switch msg.String() {
		case "down", "j":
			if p.selected < len(p.items)-1 {
				p.selected++

				// 🔥 scroll down when cursor passes viewport
				if p.selected >= p.scrollOffset+p.visibleItems() {
					p.scrollOffset++
				}
			}

		case "up", "k":
			if p.selected > 0 {
				p.selected--

				// 🔥 scroll up when cursor goes above viewport
				if p.selected < p.scrollOffset {
					p.scrollOffset--
				}
			}

		case "p":
			if len(p.items) > 0 {
				selected := p.items[p.selected]
				return p, func() tea.Msg {
					return PlaySongMsg{URI: selected.URI}
				}
			}
		}
	}

	return p, nil
}

func (p Result) View() string {
	borderColor := lipgloss.Color("#1a1a1a")
	if p.focused {
		borderColor = lipgloss.Color("#1DB954")
	}

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#1DB954")).
		PaddingLeft(1)

	selectedTitleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#1DB954")).
		Bold(true)

	normalTitleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#e0e0e0"))

	artistStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#555555"))

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
			Render("press / to search")
		return style.Render(empty)
	}

	start := p.scrollOffset
	end := start + p.visibleItems()
	if end > len(p.items) {
		end = len(p.items)
	}

	var content string
	content += titleStyle.Render("Results") + "\n\n"

	for i := start; i < end; i++ {
		item := p.items[i]
		if i == p.selected {
			line := selectedTitleStyle.Render(item.Title) + "\n" +
				artistStyle.Render(item.Artist)
			content += selectedRowStyle.Render(line) + "\n\n"
		} else {
			line := normalTitleStyle.Render(item.Title) + "\n" +
				artistStyle.Render(item.Artist)
			content += normalRowStyle.Render(line) + "\n\n"
		}
	}

	return style.Render(content)
}

func (p Result) SetResults(search models.SearchResult) Result {
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
	p.scrollOffset = 0
	return p
}

func (p Result) visibleItems() int {
	items := (p.height - 4) / 3 // 3 lines per item (title + artist + spacing)
	if items < 1 {
		return 1
	}
	return items
}
