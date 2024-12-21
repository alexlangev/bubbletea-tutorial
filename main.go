package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Where the app state is saved
type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our current cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

// returns the initial app state
func initialModel() model {
	return model{
		// Basic todo list
		choices: []string{"Buy carrots", "Buy ham", "Buy gum"},
		// keys refer to the index of the choices slice above
		selected: make(map[int]struct{}),
	}
}

// returns an initial command. Kick start the app loop?
func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("bubbletea tutorial")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a keypress?
	case tea.KeyMsg:

		// What key is actually pressed?
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			{
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}
			}

		case "enter", " ":
			{
				_, ok := m.selected[m.cursor]
				if ok {
					delete(m.selected, m.cursor)
				} else {
					m.selected[m.cursor] = struct{}{}
				}
			}
		}

	}

	// Return the update model to Bubble tea runtime for processing
	return m, nil
}

// Render the UI?
func (m model) View() string {
	// header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at it?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is the choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas,, there's been an error: %v", err)
		os.Exit(1)
	}
}
