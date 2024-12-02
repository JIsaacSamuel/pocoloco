package main

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/JIsaacSamuel/pocoloco/internal/header"
	"github.com/JIsaacSamuel/pocoloco/internal/helpers"
	nav "github.com/JIsaacSamuel/pocoloco/pkg/navigation"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("99"))

type model struct {
	hover int
	table []fs.DirEntry
}

func (m model) Init() tea.Cmd { return nil }

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
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

		case "ctrl+z":
			helpers.Go_to("..")
			m.table = nav.Get_dirs()
			m.hover = 0

		case "enter":
			if m.hover >= 0 {
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
		}
	}
	return m, nil
}

func (m model) View() string {
	s := header.Get_header()
	var fname string

	if len(m.table) == 0 {
		return fmt.Sprintf("%s %s\n", "!!", "This directory is empty.")
	}

	for i := 0; i < len(m.table); i++ {
		if i == m.hover {
			if m.table[i].IsDir() == false {
				fname = fmt.Sprintf("%s %s\n", "x", baseStyle.Render(m.table[i].Name()))
			} else {
				fname = fmt.Sprintf("%s %s\n", ">", baseStyle.Render(m.table[i].Name()))
			}
		} else {
			fname = fmt.Sprintf("%s %s\n", " ", m.table[i].Name())
		}
		s += fname
	}

	return s
}

func initialModel() *model {
	return &model{
		hover: 0,
		table: nav.Get_dirs(),
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
