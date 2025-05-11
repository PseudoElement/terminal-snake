package exitinfo

import "github.com/charmbracelet/lipgloss"

var exitInfo = lipgloss.NewStyle().
	Padding(0, 1).
	Background(lipgloss.Color("#5783f2")).
	MarginBottom(1).
	Align(lipgloss.Center, lipgloss.Center)
