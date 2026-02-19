package main

import (
	"fmt"

	"github.com/abhinav0411/spctl/spotify"
)

func main() {
	code_verifier, code := spotify.Login()
	spotify.RequestAccessToken(code_verifier, code)
	fmt.Println("Hope this works")
}
