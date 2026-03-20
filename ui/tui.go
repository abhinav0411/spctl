package ui

import "github.com/abhinav0411/spctl/models"

type spctl struct {
	currentScreen string
	client        *models.Client
	loginModel    login
}
