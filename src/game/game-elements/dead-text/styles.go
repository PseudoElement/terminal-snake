package dead_text

import "github.com/charmbracelet/lipgloss"

var deadText = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#fc0303")).
	MarginBottom(1).
	MarginLeft(4).
	Align(lipgloss.Center, lipgloss.Center).
	Width(15)
