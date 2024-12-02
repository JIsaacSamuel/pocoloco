package main

import (
	"fmt"
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
	table []string
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
				helpers.Go_to(m.table[m.hover])
				m.table = nav.Get_dirs()
			} else {
				return m, nil
			}

			if len(m.table) > 0 {
				m.hover = 0
			} else {
				m.table = append(m.table, "No files here")
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
	for i := 0; i < len(m.table); i++ {
		if i == m.hover {
			fname = fmt.Sprintf("%s %s\n", ">", baseStyle.Render(m.table[i]))
		} else {
			fname = fmt.Sprintf("%s %s\n", " ", m.table[i])
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
