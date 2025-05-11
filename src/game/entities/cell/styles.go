package cell

import "github.com/charmbracelet/lipgloss"

var (
	GrayCell = lipgloss.NewStyle().
			Background(lipgloss.Color("#fcfcfc")).
			Height(1).
			Width(1)

	RedCell = lipgloss.NewStyle().
		Background(lipgloss.Color("#f70505")).
		Height(1).
		Width(1)

	GreenCell = lipgloss.NewStyle().
			Background(lipgloss.Color("#03a309")).
			Height(1).
			Width(1)
)
