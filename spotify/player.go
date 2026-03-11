package spotify

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/abhinav0411/spctl/models"
)

func StartResume(c *models.Client, song *models.Song) {
	songJSON, err := json.Marshal(song)

	var songDetail string

	json.Unmarshal(songJSON, &songDetail)
	const url = "https://api.spotify.com/v1/me/player/play"
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	access_token := c.Session.AccessToken

	req.Header.Set("Authorization", "Bearer "+access_token)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)

	fmt.Println(res.StatusCode, res)
}
