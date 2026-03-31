package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Search struct {
	input   textinput.Model
	width   int
	focused bool
}

type SearchQueryMsg struct {
	Query string
}

func NewSearch() Search {
	ti := textinput.New()
	ti.Placeholder = "Search songs..."
	ti.CharLimit = 100

	return Search{
		input:   ti,
		focused: false, // 🔥 NOT focused by default
	}
}

func (s Search) Init() tea.Cmd {
	return textinput.Blink
}

func (s Search) Update(msg tea.Msg) (Search, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.input.Width = s.width - 12

	case tea.KeyMsg:
		switch msg.String() {

		case "/":
			// 🔥 enter search mode
			s.focused = true
			s.input.Focus()
			return s, textinput.Blink

		case "esc":
			// 🔥 exit search mode
			s.focused = false
			s.input.Blur()
			return s, nil
		case "enter":
			if s.focused {
				query := s.input.Value()

				// exit search mode
				s.focused = false
				s.input.Blur()

				return s, func() tea.Msg {
					return SearchQueryMsg{
						Query: query,
					}
				}
			}
		}
	}

	// only update input when focused
	if s.focused {
		s.input, cmd = s.input.Update(msg)
	}

	return s, cmd
}

func (s Search) View() string {
	style := lipgloss.NewStyle().
		Width(s.width).
		Padding(0, 1)

	if s.focused {
		return style.Render("Search: " + s.input.View())
	}

	// show hint when not focused
	return style.Render("Press / to search")
}

// helper
func (s Search) Value() string {
	return s.input.Value()
}
