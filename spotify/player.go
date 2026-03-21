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

func Convert(search models.SearchResult, i int) models.Song {
	var song models.Song
	song.ContextURI = search.Tracks.Items[i].URI
	song.PositionMs = 0
	song.OffSet.Position = 0
	return song
}

func StartResume(c *models.Client, song *models.Song, id string) {
	songJSON, err := json.Marshal(song)

	if err != nil {
		log.Fatal(err)
	}

	params := url.Values{}
	params.Set("device_id", id)

	reader := strings.NewReader(string(songJSON))
	url := "https://api.spotify.com/v1/me/player/play?" + params.Encode()
	req, err := http.NewRequest("PUT", url, reader)
	if err != nil {
		log.Fatal(err)
	}

	access_token := c.Session.AccessToken

	req.Header.Set("Authorization", "Bearer "+access_token)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)

	fmt.Println(res.StatusCode)
}

func Pause(c *models.Client, id string) {
	params := url.Values{}
	params.Set("device_id", id)

	url := "https://api.spotify.com/v1/me/player/pause" + params.Encode()

	req, err := http.NewRequest("PUT", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	access_token := c.Session.AccessToken
	req.Header.Set("Authorization", "Bearer "+access_token)

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 204 {
		fmt.Println("Something went wrong")
	}
}

func SkipNext(c *models.Client, id string) {

	params := url.Values{}
	params.Set("device_id", id)

	url := "https://api.spotify.com/v1/me/player/next" + params.Encode()

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	access_token := c.Session.AccessToken
	req.Header.Set("Authorization", "Bearer "+access_token)

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 204 {
		fmt.Println("Something went wrong")
	}
}

func SkipPrev(c *models.Client, id string) {
	params := url.Values{}
	params.Set("device_id", id)

	url := "https://api.spotify.com/v1/me/player/previous" + params.Encode()

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	access_token := c.Session.AccessToken
	req.Header.Set("Authorization", "Bearer "+access_token)

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 204 {
		fmt.Println("Something went wrong")
	}
}
