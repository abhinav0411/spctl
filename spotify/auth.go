package spotify

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/pkg/browser"
	"golang.org/x/oauth2"
)

func GenerateRandomString() string {
	var code_verifier = oauth2.GenerateVerifier()
	return code_verifier
}

func CreateConf() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL: "https://accounts.spotify.com/authorize",
		},
		RedirectURL: os.Getenv("REDIRECT_URL"),
	}
	return conf
}

func Login() (string, string) {
	verifier := GenerateRandomString()

	var conf = CreateConf()

	go StartCallbackServer()

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	browser.OpenURL(url)

	code := <-AuthCodeChan
	fmt.Println("Code: ", code)
	return verifier, code
}

func RequestAccessToken(code_verifier string, code string) {
	fmt.Println("hesjkdf")
	RedirectURL := os.Getenv("REDIRECT_URL")
	ClientID := os.Getenv("CLIENT_ID")

	bodyParams := url.Values{}
	bodyParams.Set("code", code)
	bodyParams.Set("grant_type", "authorization_code")
	bodyParams.Set("client_id", ClientID)
	bodyParams.Set("redirect_uri", RedirectURL)
	bodyParams.Set("code_verifier", code_verifier)

	body := bodyParams.Encode()

	reader := strings.NewReader(body)

	url := "https://accounts.spotify.com/api/token"

	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		fmt.Println("Error in req")
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error in res")
		os.Exit(1)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error in body")
		os.Exit(1)
	}

	fmt.Println("Response: ", string(resBody))
}
