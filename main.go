package main

import (
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	username
	password
)

type login struct {
	inputs []textinput.Model
	err    error
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func initialModel() login {
	var inputs[username] = textinput.New()
	inputs[username].Focus()

}
