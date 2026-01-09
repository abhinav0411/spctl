package ui

import "github.com/charmbracelet/bubbles/list"

/* Implementing the list.Item interface for Songs */
type Song struct {
	title    string
	artist   string
	duration int
}

func (s Song) FilterValue() string {
	return s.title
}

func (s Song) Description() string {
	return s.artist
}

func (s Song) Title() string {
	return s.title
}

func (s Song) Artist() string {
	return s.artist
}

func (s Song) Duration() int {
	return s.duration
}

/* Implementing list.Item interface for Playlists */

type Playlist struct {
	name       string
	totalSongs int
}

func (p Playlist) FilterValue() string {
	return p.name
}

func (p Playlist) Description() string {
	return p.name
}

func (p Playlist) Name() string {
	return p.name
}

func (p Playlist) TotalSongs() int {
	return p.totalSongs
}

/* Now creating the list.model which we will use in view.go for playlists, queue */

type List struct {
	list list.Model
}

/* Function for Queue */
func (l *List) initQueue(listWidth, listHeight int) {
	l.list = list.New([]list.Item{}, list.NewDefaultDelegate(), listWidth, listHeight)

	l.list.Title = "Queue"
	l.list.SetItems([]list.Item{
		Song{title: "Something", artist: "me", duration: 120},
		Song{title: "Something else", artist: "me again", duration: 100},
	})
}

/* Function for Playlist */
func (l *List) initPlaylist(listWidth, listHeight int) {
	l.list = list.New([]list.Item{}, list.NewDefaultDelegate(), listWidth, listHeight)

	l.list.Title = "Playlist"
	l.list.SetItems([]list.Item{
		Playlist{name: "Playlist1", totalSongs: 3},
		Playlist{name: "Playlist2", totalSongs: 2},
	})
}

func (l List) View() string {
	return l.list.View()
}
