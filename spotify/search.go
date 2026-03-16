package spotify

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/abhinav0411/spctl/models"
)

func Search(c *models.Client) models.SearchResult {
	url_ := "https://api.spotify.com/v1/search?"

	var searchResult models.SearchResult

	params := url.Values{}
	params.Add("type", "album")
	params.Add("q", "DAMN")

	new_url := url_ + params.Encode()

	req, err := http.NewRequest("GET", new_url, nil)

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

	json.Unmarshal(body, &searchResult)

	return searchResult
}
