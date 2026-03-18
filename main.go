package main

import (
	"fmt"
	"os"
	"github.com/augustofaggion/repoman/helpers"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	projects     []string
	projectPaths []string
	cursor       int
	selected     int
	quitting     bool
	inputBuffer  string // stores numeric input
}

func initialModel() model {
	var projects []string
	var projectPaths []string
	helpers.CreateProfile(&projects, &projectPaths)
	return model{
		projects:     projects,
		projectPaths: projectPaths,
		cursor:       0,
		selected:     -1,
		quitting:     false,
		inputBuffer:  "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch key {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
			m.inputBuffer = ""
		case "down", "j":
			if m.cursor < len(m.projects)-1 {
				m.cursor++
			}
			m.inputBuffer = ""
		case "enter":
			if m.inputBuffer != "" {
				idx := 0
				fmt.Sscanf(m.inputBuffer, "%d", &idx)
				if idx > 0 && idx <= len(m.projects) {
					m.selected = idx - 1
					fmt.Printf("Selected: [%d] %s\n", idx, m.projects[m.selected])
					helpers.OpenProject(m.projectPaths[m.selected])
					m.quitting = true
					return m, tea.Quit
				}
			} else {
				m.selected = m.cursor
				fmt.Printf("Selected: [%d] %s\n", m.selected+1, m.projects[m.selected])
				helpers.OpenProject(m.projectPaths[m.selected])
				m.quitting = true
				return m, tea.Quit
			}
		default:
			// Handle numeric input
			if len(key) == 1 && key[0] >= '0' && key[0] <= '9' {
				m.inputBuffer += key
				idx := 0
				fmt.Sscanf(m.inputBuffer, "%d", &idx)
				if idx > 0 && idx <= len(m.projects) {
					m.cursor = idx - 1
				}
			} else {
				m.inputBuffer = ""
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.quitting {
		goodbyeStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF5F87"))
		s := goodbyeStyle.Render("Goodbye!\n")
		if m.selected >= 0 && m.selected < len(m.projects) {
			selectedStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFD700")).Background(lipgloss.Color("#005F5F")).Padding(0, 1)
			s += selectedStyle.Render(fmt.Sprintf("Selected: [%d] %s\n", m.selected+1, m.projects[m.selected]))
		}
		return s
	}

	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00BFFF")).Background(lipgloss.Color("#1E1E1E")).Padding(1, 2)
	selectedStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFD700")).Background(lipgloss.Color("#005F5F")).Padding(0, 1)
	normalStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#D0D0D0")).Padding(0, 1)
	instructionsStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#808080")).Italic(true)
	inputStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF87")).Bold(true)
	previewStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF87")).Background(lipgloss.Color("#222222")).Padding(0, 1)

	s := titleStyle.Render("Repoman") + "\n\n"
	for i, project := range m.projects {
		idx := fmt.Sprintf("%d", i+1)
		if m.cursor == i {
			s += selectedStyle.Render("> "+idx+". "+project) + "\n"
		} else {
			s += normalStyle.Render("  "+idx+". "+project) + "\n"
		}
	}
	s += "\n" + instructionsStyle.Render("Use ↑/↓ or j/k to navigate, Enter to open, q to quit.\nSelect by typing index and Enter.")
	if m.inputBuffer != "" {
		idx := 0
		fmt.Sscanf(m.inputBuffer, "%d", &idx)
		if idx > 0 && idx <= len(m.projects) {
			s += "\n" + previewStyle.Render(fmt.Sprintf("Selected: [%d] %s", idx, m.projects[idx-1]))
		} else {
			s += "\n" + inputStyle.Render("Selected: "+m.inputBuffer)
		}
	}
	s += "\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v\n", err)
		os.Exit(1)
	}
}
