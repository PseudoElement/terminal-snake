package game

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pseudoelement/terminal-snake/src/models"
)

// MODEL DATA

type DemoProgram struct {
	program *tea.Program
}

func NewDemoGame() DemoProgram {
	p := tea.NewProgram(
		simplePage{text: "This app is under construction", updChan: make(chan int, 100)},
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
		tea.WithFPS(30),
	)
	tea.ShowCursor()

	return DemoProgram{program: p}
}

func (this *DemoProgram) Run() error {
	_, err := this.program.Run()
	return err
}

func (this *DemoProgram) Quit() {
	this.program.Kill()
}

type simplePage struct {
	text    string
	updChan chan int
}

func (s simplePage) Init() tea.Cmd {
	return nil
}

// VIEW
func (s simplePage) View() string {
	btn := lipgloss.NewStyle().
		Background(lipgloss.Color("#888B7E")).
		MarginTop(10).
		MarginRight(2).
		// Border(lipgloss.RoundedBorder(), true).
		Height(1)

	return btn.Render("Click", "me")
}

// UPDATE
func (s simplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return s, tea.Quit
		}

	case tea.MouseMsg:
		mouseMsg := msg.(tea.MouseMsg)
		println()
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

var _ models.TerminalProgram = (*DemoProgram)(nil)
