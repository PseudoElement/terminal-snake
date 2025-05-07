package game

import (
	tea "github.com/charmbracelet/bubbletea"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/controllers"
	menu_elements "github.com/pseudoelement/terminal-snake/src/game/menu-elements"
	"github.com/pseudoelement/terminal-snake/src/models"
)

// MODEL DATA

type SnakeGameProgram struct {
	program *tea.Program
}

func NewSnakeGameProgram() SnakeGameProgram {
	p := tea.NewProgram(
		&SnakeGame{},
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
	view           string
	menuController *controllers.MenuController

	width, height int
}

func (this *SnakeGame) Init() tea.Cmd {
	firstPage := menu_elements.NewFirstPage()
	selectableElems := firstPage.SelectableElems()
	firstSelectedElem := selectableElems[0]
	firstSelectedElem.Select()

	this.menuController = controllers.NewMenuController(firstPage)

	return nil
}

// VIEW
func (this *SnakeGame) View() string {
	return this.menuController.Page().View()
}

// UPDATE
func (this *SnakeGame) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "up":
			this.menuController.SelectNext()
			return this, nil
		case "down":
			this.menuController.SelectPrev()
			return this, nil
		case "enter":
			selectedEl := this.menuController.SelectedElem()
			if selectedEl != nil {
				selectedEl.Action()
				redirectableEl, ok := selectedEl.(game_abstr.IRedirectableElement)
				if ok {
					nextPage := redirectableEl.NextPage()
					this.menuController.SetPage(nextPage)
				}
			}

			return this, nil
		case "ctrl+c":
			return this, tea.Quit
		}

	case tea.MouseMsg:
		mouseMsg := msg.(tea.MouseMsg)
		if mouseMsg.Button == tea.MouseButtonLeft && mouseMsg.Action == tea.MouseActionPress {
		}
		return this, nil

	case tea.WindowSizeMsg:
		this.width = msg.(tea.WindowSizeMsg).Width
		this.height = msg.(tea.WindowSizeMsg).Height
	}

	return this, nil
}

var _ models.TerminalProgram = (*SnakeGameProgram)(nil)
