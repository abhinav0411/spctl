package spotify

import (
	"log"
	"net/http"

	"github.com/abhinav0411/spctl/models"
)

func StartResume(c *models.Client) {
	const url = "https://api.spotify.com/v1/me/player/play"
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	access_token := c.Session.AccessToken

	req.Header.Set("Authorization", "Bearer "+access_token)
	req.Header.Set("Content-Type", "application/json")
	req.
}
