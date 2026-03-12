package spotify

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/abhinav0411/spctl/models"
)

func GetDevice(c *models.Client) ([]models.PlayerDevice, error) {
	url := "https://api.spotify.com/v1/me/player/devices"

	var result struct {
		PlayerDevices []models.PlayerDevice `json:"devices"`
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	access_token := c.Session.AccessToken

	req.Header.Set("Authorization", "Bearer "+access_token)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &result)
	return result.PlayerDevices, nil
}
