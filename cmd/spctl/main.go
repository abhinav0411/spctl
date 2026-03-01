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

	new_user := spotify.GetUser(&new_client)

	fmt.Println(new_user)

	fmt.Println("Hope this works")
}
