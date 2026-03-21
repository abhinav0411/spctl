package spotify

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/abhinav0411/spctl/models"
)

func GetCurrentSong(c *models.Client) models.CurrentSong {
	url := "https://api.spotify.com/v1/me/player/currently-playing"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}

	access_token := c.Session.AccessToken
	req.Header.Set("Authorization", "Bearer "+access_token)

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var current_song models.CurrentSong

	json.Unmarshal(body, &current_song)

	return current_song
}
