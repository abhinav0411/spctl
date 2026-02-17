package main

import (
	"fmt"

	"github.com/abhinav0411/spctl/spotify"
	"golang.org/x/oauth2"
)

func main() {
	var verifier = spotify.GenerateRandomString()
	var conf = spotify.CreateConf()

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))

	fmt.Printf("visit this url \n%v\n", url)
}
