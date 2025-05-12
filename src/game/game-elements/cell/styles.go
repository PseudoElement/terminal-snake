package cell

import "github.com/charmbracelet/lipgloss"

var (
	GrayCell = lipgloss.NewStyle().
			Background(lipgloss.Color("#aeb0af")).
			Height(1).
			Width(2)

	DarkGrayCell = lipgloss.NewStyle().
			Background(lipgloss.Color("#959696")).
			Height(1).
			Width(2)

	PinkCell = lipgloss.NewStyle().
			Background(lipgloss.Color("#f57af1")).
			Height(1).
			Width(2)

	RedCell = lipgloss.NewStyle().
		Background(lipgloss.Color("#f70505")).
		Height(1).
		Width(2)

	BlueCell = lipgloss.NewStyle().
			Background(lipgloss.Color("#0335fc")).
			Height(1).
			Width(2)
)
