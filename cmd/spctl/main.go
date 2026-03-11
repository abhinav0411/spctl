package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/abhinav0411/spctl/models"
	"github.com/abhinav0411/spctl/spotify"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	code_verifier, code := spotify.Login()
	spotify.RequestToken(code_verifier, code)

	configDir, err := os.UserConfigDir()

	new_session := models.Load(configDir + "/spctl")
	new_http_client := http.Client{}

	new_client := models.Client{
		Session:    &new_session,
		HTTPClient: &new_http_client,
	}

	var song models.Song

	song.ContextURI = "spotify:track:66nBMj6cwE9V6LxcpQQpFs"
	song.OffSet.Position = 0
	song.OffSet.PositionMs = 0

	spotify.StartResume(&new_client, &song)

	fmt.Println("Hope this works")
}
