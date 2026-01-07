package main

import (
	"log"

	"github.com/abhinav0411/spctl/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := ui.NewScreen()
	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}
