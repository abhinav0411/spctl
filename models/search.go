package models

type SearchResult struct {
	Tracks struct {
		Href     string `json:"href"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
		Items    []struct {
			Album struct {
				AlbumType        string   `json:"album_type"`
				TotalTracks      int      `json:"total_tracks"`
				AvailableMarkets []string `json:"available_markets"`
				ExternalUrls     struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href   string `json:"href"`
				ID     string `json:"id"`
				Images []struct {
					URL    string `json:"url"`
					Height int    `json:"height"`
					Width  int    `json:"width"`
				} `json:"images"`
				Name                 string `json:"name"`
				ReleaseDate          string `json:"release_date"`
				ReleaseDatePrecision string `json:"release_date_precision"`
				Restrictions         struct {
					Reason string `json:"reason"`
				} `json:"restrictions"`
				Type    string `json:"type"`
				URI     string `json:"uri"`
				Artists []struct {
					ExternalUrls struct {
						Spotify string `json:"spotify"`
					} `json:"external_urls"`
					Href string `json:"href"`
					ID   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
					URI  string `json:"uri"`
				} `json:"artists"`
			} `json:"album"`
			Artists []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			AvailableMarkets []string `json:"available_markets"`
			DiscNumber       int      `json:"disc_number"`
			DurationMs       int      `json:"duration_ms"`
			Explicit         bool     `json:"explicit"`
			ExternalIds      struct {
				Isrc string `json:"isrc"`
				Ean  string `json:"ean"`
				Upc  string `json:"upc"`
			} `json:"external_ids"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href       string `json:"href"`
			ID         string `json:"id"`
			IsPlayable bool   `json:"is_playable"`
			LinkedFrom struct {
			} `json:"linked_from"`
			Restrictions struct {
				Reason string `json:"reason"`
			} `json:"restrictions"`
			Name        string `json:"name"`
			Popularity  int    `json:"popularity"`
			PreviewURL  string `json:"preview_url"`
			TrackNumber int    `json:"track_number"`
			Type        string `json:"type"`
			URI         string `json:"uri"`
			IsLocal     bool   `json:"is_local"`
		} `json:"items"`
	} `json:"tracks"`
	Albums struct {
		Href     string `json:"href"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
		Items    []struct {
			AlbumType        string   `json:"album_type"`
			TotalTracks      int      `json:"total_tracks"`
			AvailableMarkets []string `json:"available_markets"`
			ExternalUrls     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				URL    string `json:"url"`
				Height int    `json:"height"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                 string `json:"name"`
			ReleaseDate          string `json:"release_date"`
			ReleaseDatePrecision string `json:"release_date_precision"`
			Restrictions         struct {
				Reason string `json:"reason"`
			} `json:"restrictions"`
			Type    string `json:"type"`
			URI     string `json:"uri"`
			Artists []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
		} `json:"items"`
	} `json:"albums"`
	Playlists struct {
		Href     string `json:"href"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
		Items    []struct {
			Collaborative bool   `json:"collaborative"`
			Description   string `json:"description"`
			ExternalUrls  struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				URL    string `json:"url"`
				Height int    `json:"height"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name  string `json:"name"`
			Owner struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href        string `json:"href"`
				ID          string `json:"id"`
				Type        string `json:"type"`
				URI         string `json:"uri"`
				DisplayName string `json:"display_name"`
			} `json:"owner"`
			Public     bool   `json:"public"`
			SnapshotID string `json:"snapshot_id"`
			Items      struct {
				Href  string `json:"href"`
				Total int    `json:"total"`
			} `json:"items"`
			Tracks struct {
				Href  string `json:"href"`
				Total int    `json:"total"`
			} `json:"tracks"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"items"`
	} `json:"playlists"`
}
