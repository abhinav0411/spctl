package spotify

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/abhinav0411/spctl/models"
)

func ConvertSearch(search models.SearchResult, i int) models.Song {
	var song models.Song
	song.ContextURI = search.Tracks.Items[i].URI
	song.PositionMs = 0
	song.OffSet.Position = 0
	return song
}

func ConvertResume(current_song models.CurrentSong) models.Song {
	var song models.Song
	song.ContextURI = current_song.Item.URI
	song.PositionMs = current_song.ProgressMs
	song.OffSet.Position = 0
	return song
}

func TransferPlayback(c *models.Client, deviceID string) {
	body := fmt.Sprintf(`{"device_ids": ["%s"], "play": false}`, deviceID)

	req, _ := http.NewRequest(
		"PUT",
		"https://api.spotify.com/v1/me/player",
		strings.NewReader(body),
	)

	req.Header.Set("Authorization", "Bearer "+c.Session.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	c.HTTPClient.Do(req)
}

func Start(c *models.Client, body interface{}, id string) {
	songJSON, _ := json.Marshal(body)

	params := url.Values{}
	params.Set("device_id", id)

	reader := strings.NewReader(string(songJSON))
	url := "https://api.spotify.com/v1/me/player/play?" + params.Encode()

	req, _ := http.NewRequest("PUT", url, reader)

	req.Header.Set("Authorization", "Bearer "+c.Session.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	c.HTTPClient.Do(req)
}

func Resume(c *models.Client, id string) {
	params := url.Values{}
	params.Set("device_id", id)

	url := "https://api.spotify.com/v1/me/player/play?" + params.Encode()

	req, _ := http.NewRequest("PUT", url, nil)

	req.Header.Set("Authorization", "Bearer "+c.Session.AccessToken)

	c.HTTPClient.Do(req)
}

func Pause(c *models.Client, id string) {
	params := url.Values{}
	params.Set("device_id", id)

	url := "https://api.spotify.com/v1/me/player/pause?" + params.Encode()

	req, err := http.NewRequest("PUT", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	access_token := c.Session.AccessToken
	req.Header.Set("Authorization", "Bearer "+access_token)

	_, err = c.HTTPClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}
}

func SkipNext(c *models.Client, id string) {

	params := url.Values{}
	params.Set("device_id", id)

	url := "https://api.spotify.com/v1/me/player/next?" + params.Encode()

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	access_token := c.Session.AccessToken
	req.Header.Set("Authorization", "Bearer "+access_token)

	_, err = c.HTTPClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}
}

func SkipPrev(c *models.Client, id string) {
	params := url.Values{}
	params.Set("device_id", id)

	url := "https://api.spotify.com/v1/me/player/previous?" + params.Encode()

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	access_token := c.Session.AccessToken
	req.Header.Set("Authorization", "Bearer "+access_token)

	_, err = c.HTTPClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}
}
