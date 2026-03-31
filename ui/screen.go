package ui

import (
	"github.com/abhinav0411/spctl/models"
	"github.com/abhinav0411/spctl/spotify"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type screen struct {
	width  int
	height int

	player Player
	search Search
	result Playlist

	logo string
}

func NewScreen() *screen {
	return &screen{
		player: NewPlayer(),
		search: NewSearch(),
		result: NewPlaylist(),
		logo: `
                             █████    ████ 
                            ░░███    ░░███ 
  █████  ████████   ██████  ███████   ░███ 
 ███░░  ░░███░░███ ███░░███░░░███░    ░███ 
░░█████  ░███ ░███░███ ░░░   ░███     ░███ 
 ░░░░███ ░███ ░███░███  ███  ░███ ███ ░███ 
 ██████  ░███████ ░░██████   ░░█████  █████
░░░░░░   ░███░░░   ░░░░░░     ░░░░░  ░░░░░ 
         ░███                              
         █████                             
        ░░░░░                              
`,
	}
}

func (s *screen) ScreenUpdate(
	msg tea.Msg,
	song models.CurrentSong,
	client models.Client,
	device models.PlayerDevice,
) tea.Cmd {

	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height

	case models.SearchResult:
		// update results (no printing!)
		s.result = s.result.SetResults(msg)

	case PlaySongMsg:
		cmds = append(cmds, playSongCmd(client, msg.URI, device.ID))

	case SearchQueryMsg:
		cmds = append(cmds, searchCmd(&client, msg.Query))
	}

	// --- SEARCH ---
	var searchUpdateCmd tea.Cmd
	s.search, searchUpdateCmd = s.search.Update(msg)
	cmds = append(cmds, searchUpdateCmd)

	// --- PLAYLIST ---
	var playlistCmd tea.Cmd
	if s.search.focused {
		s.result, playlistCmd = s.result.Update(nil)
	} else {
		s.result, playlistCmd = s.result.Update(msg)
	}
	cmds = append(cmds, playlistCmd)

	// --- PLAYER ---
	s.player = s.player.SetSong(song)

	var playerCmd tea.Cmd

	if s.search.focused {
		s.player, playerCmd = s.player.Update(nil, &client, device.ID)

	} else if key, ok := msg.(tea.KeyMsg); ok {

		switch key.String() {
		case "up", "down", "j", "k", "enter":
			// playlist owns these
			s.player, playerCmd = s.player.Update(nil, &client, device.ID)

		default:
			s.player, playerCmd = s.player.Update(msg, &client, device.ID)
		}

	} else {
		s.player, playerCmd = s.player.Update(msg, &client, device.ID)
	}

	cmds = append(cmds, playerCmd)

	return tea.Batch(cmds...)
}

func (s *screen) ScreenView() string {
	if s.width == 0 || s.height == 0 {
		return "Initializing..."
	}

	// --- TOP ---
	top := s.search.View()

	// --- LEFT ---
	left := s.result.View()

	// --- CENTER ---
	center := lipgloss.NewStyle().
		Width(s.width/3).
		Align(lipgloss.Center, lipgloss.Center).
		Render(s.logo)

	// --- RIGHT ---
	right := lipgloss.NewStyle().
		Width(s.width / 3).
		Render("")

	middle := lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		center,
		right,
	)

	// --- BOTTOM ---
	bottom := s.player.View()

	return lipgloss.JoinVertical(
		lipgloss.Top,
		top,
		middle,
		bottom,
	)
}

// --- COMMANDS ---

func playSongCmd(client models.Client, uri string, deviceID string) tea.Cmd {
	return func() tea.Msg {

		body := map[string]interface{}{
			"uris": []string{uri},
		}

		spotify.Start(&client, body, deviceID)

		return nil
	}
}

func searchCmd(client *models.Client, query string) tea.Cmd {
	return func() tea.Msg {
		return spotify.Search(client, query)
	}
}
