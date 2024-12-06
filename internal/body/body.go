package body

import (
	"fmt"
	"io/fs"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/sys/unix"
)

func getTerminalSize() (int, error) {
	var ws *unix.Winsize

	// Use ioctl syscall to get terminal window size
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
	if err != nil {
		return 0, err
	}

	return int(ws.Row), nil
}

var (
	textStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("231"))

	baseStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("46"))
)

func Body(table []fs.DirEntry, hover, index int) string {
	var s string
	var prefix string
	var suffix string
	var style lipgloss.Style

	lines, err := getTerminalSize()
	if err != nil || lines-13 <= 0 {
		lines = 10
	} else {
		lines -= 13
	}

	pointer := ""

	if len(table) == 0 {
		s += fmt.Sprintf("%s %s\n", "!!", "Nothing matches your search.")
	}

	if hover > index+(lines) {
		index = hover - lines
	}

	for i := index; i < min(index+lines+1, len(table)); i++ {
		if i == hover {
			style = baseStyle
			if !table[i].IsDir() {
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

	if index == 0 {
		prefix = "-- Top --\n"
	} else {
		prefix = "...\n"
	}

	if index+lines >= len(table)-1 {
		suffix = "-- End --\n"
	} else {
		suffix = "...\n "
	}

	return prefix + s + suffix
}
