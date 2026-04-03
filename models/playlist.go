package models

type PlaylistResponse struct {
	Items []Playlist `json:"items"`
}

type Playlist struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	URI  string `json:"uri"`
}

type PlaylistTracksResponse struct {
	Items []PlaylistTrackItem `json:"items"`
	Total int                 `json:"total"`
}

type PlaylistTrackItem struct {
	Track *PlaylistTrack `json:"track"`
	Item  *PlaylistTrack `json:"item"` // some entries use "item" instead
}

type PlaylistTrack struct {
	Name    string           `json:"name"`
	URI     string           `json:"uri"`
	Artists []PlaylistArtist `json:"artists"`
}

type PlaylistArtist struct {
	Name string `json:"name"`
}
