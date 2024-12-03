package body

import (
	"fmt"
	"io/fs"

	"github.com/charmbracelet/lipgloss"
)

var textStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("231"))

var baseStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("46"))

func Body(table []fs.DirEntry, hover, index int) string {
	var s string
	pointer := ""
	style := textStyle

	if len(table) == 0 {
		s += fmt.Sprintf("%s %s\n", "!!", "This directory is empty.")
	}

	if hover > index+9 {
		index = hover - 9
	}

	for i := index; i < min(index+10, len(table)); i++ {
		if i == hover {
			style = baseStyle
			if table[i].IsDir() == false {
				pointer = "x"
			} else {
				pointer = ">"
			}
		} else {
			pointer = " "
			style = textStyle
		}
		s += fmt.Sprintf("%s %s\n", pointer, style.Render(table[i].Name()))
	}

	return s
}
