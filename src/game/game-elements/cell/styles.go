package cell

import "github.com/charmbracelet/lipgloss"

var (
	GrayCell = lipgloss.NewStyle().
			Background(lipgloss.Color("#aeb0af")).
			Border(lipgloss.NormalBorder(), false).
			Height(1).
			Width(2)
	DarkGrayCell = lipgloss.NewStyle().
			Background(lipgloss.Color("#959696")).
			Border(lipgloss.NormalBorder(), false).
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
