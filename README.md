# Pocoloco

Pocoloco is a terminal-based application built in Go using the [Bubble Tea](https://github.com/charmbracelet/bubbletea) package for easier navigation of terminal.

## Features
- **File Navigation:** Navigate through directories using intuitive keyboard shortcuts.
- **File Editing:** Open and edit files with `nano` directly from the terminal.
- **Search and Filter:** Search and filter files dynamically as you type.

## Pre-requisites
- Make sure you are on linux or desbian based system. This files does not run on Windows, however is compatible with WSL.
- Make sure you have the `xclip` along with `nano` and `code`
- You can download `xclip` by running the following command:
  ```bash
  sudo apt-get install xclip

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/JIsaacSamuel/pocoloco.git

2. Navigate into the directory where `main.go` resides:
   ```bash
   cd cmd/pocoloco

3. Run the Go project:
   ```bash
   go run main.go

This will open the command in an alternate buffer screen within the terminal. If you are using WSL2 with ubuntu distro like me you can access the command `pocoloco` anywhere in the terminal by following steps given below:
1. Build the executable:
    ```bash 
    go build
2. Move the executable to `/usr/local/bin/`:
    ```bash
    mv pocoloco /usr/local/bin/

## Key Commands

| Key           | Action                                            |
|---------------|---------------------------------------------------|
| `up`          | Move the selection cursor up in the list.         |
| `down`        | Move the selection cursor down in the list.       |
| `ctrl+c`      | Copy the path of current directory and close the application. |
| `esc`         | Clear the screen and quit the application.        |
| `enter`       | Open the selected file or directory.              |
| `ctrl+z`      | Navigate to the parent directory.                 |
| `ctrl+s`      | Executes `code .` in the current directory.       |
| `backspace`   | Delete the last character of the search query.    |
| Dynamic Input | Filter files and directories dynamically by typing. |
