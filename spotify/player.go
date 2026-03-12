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

func StartResume(c *models.Client, song *models.Song) {
	songJSON, err := json.Marshal(song)

	if err != nil {
		log.Fatal(err)
	}

	var device []models.PlayerDevice
	device, err = GetDevice(c)

	if err != nil {
		log.Fatal(err)
	}
	device_id := device[0].ID
	fmt.Println(device)

	params := url.Values{}
	params.Set("device_id", device_id)
	fmt.Println(device_id)

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
