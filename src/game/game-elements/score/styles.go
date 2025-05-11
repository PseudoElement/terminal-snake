package score

import "github.com/charmbracelet/lipgloss"

var scoreTab = lipgloss.NewStyle().
	Padding(0, 1).
	Border(lipgloss.RoundedBorder(), true).
	Align(lipgloss.Center, lipgloss.Center)
