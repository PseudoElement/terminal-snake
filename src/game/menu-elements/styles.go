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

var (
	textRules = lipgloss.NewStyle().
		Background(lipgloss.Color("#d9d5c1")).
		Foreground(lipgloss.Color("#078c3f")).
		Padding(0, 2).
		MarginBottom(1).
		Align(lipgloss.Left, lipgloss.Center)
)
