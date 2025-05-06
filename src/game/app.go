package game

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	menu_elements "github.com/pseudoelement/terminal-snake/src/game/menu-elements"
	"github.com/pseudoelement/terminal-snake/src/models"
)

// MODEL DATA

type SnakeGameProgram struct {
	program *tea.Program
}

func NewSnakeGameProgram() SnakeGameProgram {
	p := tea.NewProgram(
		SnakeGame{},
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
		tea.WithFPS(30),
	)
	tea.ShowCursor()

	return SnakeGameProgram{program: p}
}

func (this *SnakeGameProgram) Run() error {
	_, err := this.program.Run()
	return err
}

func (this *SnakeGameProgram) Quit() {
	this.program.Kill()
}

type SnakeGame struct {
	viewElements menu_elements.ViewElements
}

func (s SnakeGame) Init() tea.Cmd {
	return nil
}

// VIEW
func (s SnakeGame) View() string {
	btn := lipgloss.NewStyle().
		Background(lipgloss.Color("#888B7E")).
		MarginTop(10).
		MarginRight(2).
		// Border(lipgloss.RoundedBorder(), true).
		Height(1)

	return btn.Render("Click", "me")
}

// UPDATE
func (s SnakeGame) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return s, tea.Quit
		}

	case tea.MouseMsg:
		mouseMsg := msg.(tea.MouseMsg)
		if mouseMsg.Button == tea.MouseButtonLeft && mouseMsg.Action == tea.MouseActionPress {
			fmt.Println("Hello world")
		}
		return s, nil

	case tea.WindowSizeMsg:
		return s, func() tea.Msg {
			return msg
		}
	}

	return s, nil
}

var _ models.TerminalProgram = (*SnakeGameProgram)(nil)
