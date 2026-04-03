package ui

import (
	"github.com/abhinav0411/spctl/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	PanelResult = iota
	PanelPlaylist
)

type screen struct {
	width  int
	height int

	player   Player
	search   Search
	result   Result
	playlist Playlist

	logo            string
	playlistsLoaded bool
	activePanel     int
}

func NewScreen() *screen {
	s := &screen{
		player:      NewPlayer(),
		search:      NewSearch(),
		result:      NewResult(),
		playlist:    NewPlaylist(),
		activePanel: PanelResult,
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
	// Result panel is focused by default
	s.result.focused = true
	return s
}

func (s *screen) cycleFocus() {
	// Only cycle between result and playlist (not search — search has its own toggle)
	if s.activePanel == PanelResult {
		s.activePanel = PanelPlaylist
		s.result.focused = false
		s.playlist.focused = true
	} else {
		s.activePanel = PanelResult
		s.result.focused = true
		s.playlist.focused = false
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

		if !s.playlistsLoaded {
			cmds = append(cmds, fetchPlaylistsCmd(&client))
			s.playlistsLoaded = true
		}

	case tea.KeyMsg:
		if msg.String() == "tab" && !s.search.focused {
			s.cycleFocus()
			return tea.Batch(cmds...)
		}
		// add this — if / is pressed, let search handle it exclusively
		if msg.String() == "/" && !s.search.focused {
			s.search, _ = s.search.Update(msg)
			return tea.Batch(cmds...)
		}

	// 🔍 SEARCH RESULTS
	case models.SearchResult:
		s.result = s.result.SetResults(msg)
		s.activePanel = PanelResult
		s.result.focused = true
		s.playlist.focused = false
		// also clear any playlist view state
		s.result.scrollOffset = 0
		s.result.selected = 0

	// 🎧 PLAYLIST LIST
	case models.PlaylistResponse:
		s.playlist = s.playlist.SetPlaylists(msg)

	// 🎵 PLAYLIST TRACKS → reuse result panel
	case models.PlaylistTracksResponse:
		var items []TrackItem
		for _, item := range msg.Items {
			track := item.Track
			if track == nil {
				track = item.Item
			}
			if track == nil || track.Name == "" {
				continue
			}
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

		s.result.items = items
		s.result.selected = 0
		s.result.scrollOffset = 0
		// Switch focus to result panel so the user can browse the loaded tracks
		s.activePanel = PanelResult
		s.result.focused = true
		s.playlist.focused = false

	// ▶️ PLAY SONG
	case PlaySongMsg:
		cmds = append(cmds, playSongCmd(client, msg.URI, device.ID))

	case SelectPlaylistMsg:
		if msg.Play {
			cmds = append(cmds, playPlaylistCmd(client, msg.URI, device.ID))
			cmds = append(cmds, fetchPlaylistTracksCmd(&client, msg.ID))
		} else {
			cmds = append(cmds, fetchPlaylistTracksCmd(&client, msg.ID))
		}

	// 🔍 SEARCH QUERY
	case SearchQueryMsg:
		cmds = append(cmds, searchCmd(&client, msg.Query))
	}

	// --- SEARCH ---
	var searchCmd tea.Cmd
	s.search, searchCmd = s.search.Update(msg)
	cmds = append(cmds, searchCmd)

	// --- RESULT (LEFT PANEL) ---
	// Block input to result when search bar is focused
	// --- RESULT (LEFT PANEL) ---
	var resultCmd tea.Cmd
	if s.search.focused || s.activePanel != PanelResult {
		if _, ok := msg.(tea.WindowSizeMsg); ok {
			s.result, resultCmd = s.result.Update(msg)
		} else if _, ok := msg.(models.SearchResult); ok {
			s.result, resultCmd = s.result.Update(msg) // always pass search results
		} else {
			s.result, resultCmd = s.result.Update(nil)
		}
	} else {
		s.result, resultCmd = s.result.Update(msg)
	}
	cmds = append(cmds, resultCmd)

	// --- PLAYLIST (RIGHT PANEL) ---
	var playlistCmd tea.Cmd
	if s.search.focused || s.activePanel != PanelPlaylist {
		if _, ok := msg.(tea.WindowSizeMsg); ok {
			s.playlist, playlistCmd = s.playlist.Update(msg) // always pass size
		} else {
			s.playlist, playlistCmd = s.playlist.Update(nil)
		}
	} else {
		s.playlist, playlistCmd = s.playlist.Update(msg)
	}
	cmds = append(cmds, playlistCmd)

	// --- PLAYER ---
	s.player = s.player.SetSong(song)

	var playerCmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if s.search.focused {
			// Search is typing — block player navigation
			s.player, playerCmd = s.player.Update(nil, &client, device.ID)

		} else if s.result.focused || s.playlist.focused {
			switch msg.String() {
			case "up", "down", "j", "k", "enter":
				s.player, playerCmd = s.player.Update(nil, &client, device.ID)
			default:
				s.player, playerCmd = s.player.Update(msg, &client, device.ID)
			}
		} else {
			s.player, playerCmd = s.player.Update(msg, &client, device.ID)
		}

	default:
		s.player, playerCmd = s.player.Update(msg, &client, device.ID)
	}

	cmds = append(cmds, playerCmd)

	return tea.Batch(cmds...)
}

func (s *screen) ScreenView() string {
	if s.width == 0 || s.height == 0 {
		return "Initializing..."
	}

	// heights
	searchHeight := lipgloss.Height(s.search.View())
	playerHeight := lipgloss.Height(s.player.View())
	middleHeight := s.height - searchHeight - playerHeight

	// --- TOP ---
	top := lipgloss.NewStyle().
		Width(s.width).
		Render(s.search.View())

	// --- CENTER LOGO + KEYBINDS ---
	logoStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#1DB954")).
		Bold(true).
		Align(lipgloss.Center)

	keyStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#444444"))

	accentKey := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Bold(true)

	keybinds := "\n\n" +
		accentKey.Render("enter") + keyStyle.Render("  play track") + "\n" +
		accentKey.Render("v    ") + keyStyle.Render("  view playlist") + "\n" +
		accentKey.Render("p    ") + keyStyle.Render("  play playlist") + "\n" +
		accentKey.Render("tab  ") + keyStyle.Render("  switch panel") + "\n" +
		accentKey.Render("space") + keyStyle.Render("  pause/resume") + "\n" +
		accentKey.Render("n / b") + keyStyle.Render("  next/prev")

	center := lipgloss.NewStyle().
		Width(s.width/3).
		Height(middleHeight).
		Align(lipgloss.Center, lipgloss.Center).
		Render(logoStyle.Render(s.logo) + keybinds)

	// --- MIDDLE ---
	left := s.result.View()
	right := s.playlist.View()

	middle := lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		center,
		right,
	)

	// --- BOTTOM ---
	bottom := lipgloss.NewStyle().
		Width(s.width).
		Render(s.player.View())

	return lipgloss.JoinVertical(
		lipgloss.Top,
		top,
		middle,
		bottom,
	)
}
