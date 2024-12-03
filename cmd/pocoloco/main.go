package main

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/JIsaacSamuel/pocoloco/internal/body"
	"github.com/JIsaacSamuel/pocoloco/internal/header"
	"github.com/JIsaacSamuel/pocoloco/internal/helpers"
	nav "github.com/JIsaacSamuel/pocoloco/pkg/navigation"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var searchQueryStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99"))

var textStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("231"))

var baseStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("46"))

type model struct {
	hover        int
	table        []fs.DirEntry
	search_query string
	win_index    int
}

func (m model) Init() tea.Cmd { return nil }

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "esc":
			m.Update(tea.ClearScreen())
			return m, tea.Quit

		case "up":
			if m.hover > 0 {
				m.hover--
			} else {
				m.hover = len(m.table) - 1
			}

		case "down":
			if m.hover < len(m.table)-1 {
				m.hover++
			} else {
				m.hover = 0
			}

		case "backspace":
			if len(m.search_query) > 0 {
				m.search_query = m.search_query[0 : len(m.search_query)-1]
			}

		case "ctrl+z":
			helpers.Go_to("..")
			m.table = nav.Get_dirs()
			m.hover = 0
			m.search_query = ""

		case "enter":
			if m.hover >= 0 {
				m.search_query = ""

				if m.table[m.hover].IsDir() == false {
					helpers.Open_nano(m.table[m.hover].Name())
					return m, tea.ClearScreen
				}

				helpers.Go_to(m.table[m.hover].Name())
				m.table = nav.Get_dirs()
			} else {
				return m, nil
			}

			if len(m.table) > 0 {
				m.hover = 0
			} else {
				m.hover = -1
			}

		case "ctrl+s":
			helpers.Start_coding()

		default:
			m.search_query += msg.String()
			m.table = helpers.Filer_files(m.table, m.search_query)
			m.hover = 0
		}
	}
	return m, nil
}

func (m model) View() string {
	head := header.Get_header()

	searchBar := searchQueryStyle.Render(m.search_query)
	searchBar += "\n"

	s := body.Body(m.table, m.hover, m.win_index)

	return head + searchBar + s
}

func initialModel() *model {
	return &model{
		hover:        0,
		table:        nav.Get_dirs(),
		search_query: "",
		win_index:    0,
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
