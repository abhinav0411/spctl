package spotify

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/abhinav0411/spctl/models"
)

func GetUser(c *models.Client) models.CurrentUser {
	const url = "https://api.spotify.com/v1/me"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	access_token := c.Session.AccessToken

	req.Header.Set("Authorization", "Bearer "+access_token)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var new_user models.CurrentUser
	json.Unmarshal(body, &new_user)

	return new_user
}
