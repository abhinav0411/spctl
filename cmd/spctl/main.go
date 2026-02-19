package main

import (
	"fmt"
	"log"

	"github.com/abhinav0411/spctl/spotify"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	code_verifier, code := spotify.Login()
	spotify.RequestAccessToken(code_verifier, code)
	fmt.Println("Hope this works")
}
