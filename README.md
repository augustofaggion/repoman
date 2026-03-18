# Repoman

## What is Repoman?
Repoman is a modern CLI tool for developers to efficiently manage and open their local repositories. It provides a simple, interactive terminal interface using Bubble Tea and Lipgloss, allowing you to quickly browse, select, and open projects from a specified directory.

## How to Use

### Quick Install (Recommended)

1. **Install Repoman with Go**
   ```sh
   go install github.com/augustofaggion/repoman@latest
   ```
   (The module path in go.mod is set to github.com/augustofaggion/repoman.)

   **If you fork this repo and want to use your own GitHub username, update the module path in go.mod to match your fork’s GitHub URL.**

2. **Run Repoman**
   - Simply type:
     ```sh
     repoman
     ```

### Manual Setup

1. **Clone the repository**
   - Ensure you have Go installed.
   - Clone the repository and run `go mod tidy` to install dependencies.

2. **Build and run manually**
   - Build:
     ```sh
     go build -o repoman main.go
     ```
   - Run:
     ```sh
     ./repoman
     ```

---

**First Launch**
   - On first run, you will be prompted to enter your main repo directory path. This will be saved in `profile.json`.

**Project Selection**
   - Use the arrow keys or `j/k` to navigate the list.
   - Type a project index number to preview and select.
   - Press `Enter` to open the selected project in VS Code.
   - Press `q` to quit.

## Features
- Interactive TUI for project selection
- Instant preview by typing project index
- Opens projects in VS Code
- Remembers your repo directory

## Requirements
- Go 1.18+
- VS Code (for project opening)

## License
MIT
