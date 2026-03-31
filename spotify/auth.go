package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/abhinav0411/spctl/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"github.com/pkg/browser"
	"golang.org/x/oauth2"
)

func GenerateRandomString() string {
	var code_verifier = oauth2.GenerateVerifier()
	return code_verifier
}

func CreateConf() *oauth2.Config {
	scope_list := []string{"user-read-playback-state", "user-modify-playback-state", "user-read-currently-playing", "streaming", "playlist-read-private", "user-follow-read", "user-read-playback-position", "user-top-read", "user-read-recently-played", "user-library-read", "user-read-private"}
	conf := &oauth2.Config{
		ClientID: os.Getenv("CLIENT_ID"),
		Endpoint: oauth2.Endpoint{
			AuthURL: "https://accounts.spotify.com/authorize",
		},
		RedirectURL: os.Getenv("REDIRECT_URL"),
		Scopes:      scope_list,
	}
	return conf
}

func Createdir() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error while getting the config dir")
		os.Exit(1)
	}

	path := dir + "/" + "spctl"
	_, err = os.Stat(path)
	if err == nil {
		return path
	}
	err = os.Mkdir(path, 0o777)

	if err != nil {
		fmt.Println("Error while creating the dir")
		os.Exit(1)
	}
	return path
}

func Login() (string, string) {

	verifier := GenerateRandomString()

	var conf = CreateConf()

	go StartCallbackServer()

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	browser.OpenURL(url)

	code := <-AuthCodeChan
	return verifier, code
}

func LoginCmd() (tea.Msg, models.Client) {
	err := godotenv.Load()
	if err != nil {
		return "ERROR error loading env.", models.Client{}
	}
	code_Verifier, code := Login()
	RequestToken(code_Verifier, code)
	configDir, err := os.UserConfigDir()
	new_session := models.Load(configDir + "/spctl")
	new_http_client := http.Client{}

	new_client := models.Client{
		Session:    &new_session,
		HTTPClient: &new_http_client,
	}

	return "Login complete", new_client
}

func RequestToken(code_verifier string, code string) {
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

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println("Error in req")
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error in res", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error in body")
		os.Exit(1)
	}

	// Creating the directory and the json file
	path := Createdir()

	var new_session models.Session
	json.Unmarshal(resBody, &new_session)

	new_session.Save(path)
}

func NewAccessToken(c *models.Client, session *models.Session) {
	refresh_token := c.Session.RefreshToken
	client_id := os.Getenv("CLIENT_ID")

	url_ := "https://accounts.spotify.com/api/token"

	bodyParams := url.Values{}
	bodyParams.Set("client_id", client_id)
	bodyParams.Set("grant_type", "refresh_token")
	bodyParams.Set("refresh_token", refresh_token)

	body := bodyParams.Encode()
	reader := strings.NewReader(body)

	req, err := http.NewRequest("POST", url_, reader)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.HTTPClient.Do(req)

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

	json.Unmarshal(resBody, &session)

	dir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error while getting the config dir")
		os.Exit(1)
	}

	path := dir + "/" + "spctl"
	session.Save(path)
}
