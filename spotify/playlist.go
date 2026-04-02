package spotify

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/abhinav0411/spctl/models"
)

func GetUserPlaylists(c *models.Client) models.PlaylistResponse {
	url_ := "https://api.spotify.com/v1/me/playlists"

	req, err := http.NewRequest("GET", url_, nil)
	if err != nil {
		log.Fatal(err)
	}

	params := url.Values{}
	params.Add("limit", "5")

	access_token := c.Session.AccessToken
	req.Header.Set("Authorization", "Bearer "+access_token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var playlists models.PlaylistResponse
	json.Unmarshal(body, &playlists)

	return playlists
}

func GetPlaylistTracks(c *models.Client, playlistID string) models.PlaylistTracksResponse {
	baseURL := "https://api.spotify.com/v1/playlists/" + playlistID + "/items"

	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("limit", "50")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", "Bearer "+c.Session.AccessToken)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var tracks models.PlaylistTracksResponse
	json.Unmarshal(body, &tracks)
	return tracks
}

func PlayPlaylist(c *models.Client, playlistURI string, deviceID string) {
	url := "https://api.spotify.com/v1/me/player/play"

	bodyData := map[string]string{
		"context_uri": playlistURI,
	}

	jsonData, _ := json.Marshal(bodyData)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	access_token := c.Session.AccessToken
	req.Header.Set("Authorization", "Bearer "+access_token)
	req.Header.Set("Content-Type", "application/json")

	// optional but recommended
	if deviceID != "" {
		q := req.URL.Query()
		q.Add("device_id", deviceID)
		req.URL.RawQuery = q.Encode()
	}

	_, err = c.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}
