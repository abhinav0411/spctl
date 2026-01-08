package ui

import "github.com/charmbracelet/bubbles/list"

/* Implementing the list.Item interface for songs */
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

/* Now creating the list.model which we will use in view.go for playlists, queue */

type ListSong struct {
	list list.Model
}

func (l *ListSong) initQueue(listWidth, listHeight int) {
	l.list = list.New([]list.Item{}, list.NewDefaultDelegate(), listWidth, listHeight)

	l.list.Title = "Queue"
	l.list.SetItems([]list.Item{
		Song{title: "Something", artist: "me", duration: 120},
		Song{title: "Something else", artist: "me again", duration: 100},
	})
}

func (l *ListSong) initPlaylist(listWidth, listHeight int) {
	l.list = list.New([]list.Item{}, list.NewDefaultDelegate(), listWidth, listHeight)

	l.list.Title = "Playlist"
	l.list.SetItems([]list.Item{
		Song{title: "Playlist1", artist: "me", duration: 120},
		Song{title: "Playlist2", artist: "me again", duration: 100},
	})
}

func (l ListSong) View() string {
	return l.list.View()
}
