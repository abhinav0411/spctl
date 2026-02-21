package models

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Session struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpTime      time.Time `json:"expires_in"`
	TokenType    string    `json:"token_type"`
}

func (s Session) Save(path string) {
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("Error while marshal")
		os.Exit(1)
	}
	err = os.WriteFile(path+"/auth.json", data, 0666)
	if err != nil {
		fmt.Println("Error while saving the file")
		os.Exit(1)
	}
}

func (s Session) Load(path string) Session {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error while opening the file")
		os.Exit(1)
	}
	var new_session Session
	json.Unmarshal(file, &new_session)
	return new_session
}
