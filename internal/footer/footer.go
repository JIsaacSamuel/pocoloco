package footer

import (
	"github.com/JIsaacSamuel/pocoloco/internal/helpers"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/sys/unix"
)

var style = lipgloss.NewStyle().Foreground(lipgloss.Color("15")).Background(lipgloss.Color("30"))

func getTerminalSize() (int, error) {
	var ws *unix.Winsize

	// Use ioctl syscall to get terminal window size
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
	if err != nil {
		return 0, err
	}

	return int(ws.Col), nil
}

func Footer() string {
	cols, err := getTerminalSize()

	if err != nil {
		return style.Render(helpers.Curr_dir())
	}

	return style.Width(cols).Render(helpers.Curr_dir()) + "\n"
}
