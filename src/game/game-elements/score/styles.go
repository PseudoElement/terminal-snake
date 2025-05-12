package score

import "github.com/charmbracelet/lipgloss"

var scoreTab = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), true).
	Align(lipgloss.Center, lipgloss.Center).
	Width(15)
