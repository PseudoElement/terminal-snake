package menu_elements

import "github.com/charmbracelet/lipgloss"

var (
	selectedBtn = lipgloss.NewStyle().
			Background(lipgloss.Color("#fcba03")).
			Padding(0, 1).
			MarginBottom(1).
			Align(lipgloss.Center, lipgloss.Center).
			Height(1).
			Width(20)

	bluredBtn = lipgloss.NewStyle().
			Background(lipgloss.Color("#90918d")).
			Padding(0, 1).
			MarginBottom(1).
			Align(lipgloss.Center, lipgloss.Center).
			Height(1).
			Width(20)
)
