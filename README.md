# Pocoloco

Pocoloco is a terminal-based application built in Go using the [Bubble Tea](https://github.com/charmbracelet/bubbletea) package for creating modern, intuitive terminal user interfaces. This project aims to streamline file navigation and editing directly within the terminal.

## Features
- **File Navigation:** Navigate through directories using intuitive keyboard shortcuts.
- **File Editing:** Open and edit files with `nano` directly from the terminal.
- **Search and Filter:** Search and filter files dynamically as you type.
- **Customizable:** Easily extendable for additional features.

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/JIsaacSamuel/pocoloco.git
   cd pocoloco

## Key Commands

| Key           | Action                                            |
|---------------|---------------------------------------------------|
| `up`          | Move the selection cursor up in the list.         |
| `down`        | Move the selection cursor down in the list.       |
| `ctrl+c`      | Clear the screen and quit the application.        |
| `esc`         | Clear the screen and quit the application.        |
| `enter`       | Open the selected file or directory.              |
| `ctrl+z`      | Navigate to the parent directory.                 |
| `ctrl+s`      | Start coding with the predefined command.         |
| `backspace`   | Delete the last character of the search query.    |
| Dynamic Input | Filter files and directories dynamically by typing. |
