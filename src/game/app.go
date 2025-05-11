package game

import (
	tea "github.com/charmbracelet/bubbletea"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/controllers"
	diff_levels "github.com/pseudoelement/terminal-snake/src/game/game-elements/difficulty-levels"

	menu_elements "github.com/pseudoelement/terminal-snake/src/game/menu-elements"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	"github.com/pseudoelement/terminal-snake/src/models"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
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
	gameController *controllers.GameController
	menuController *controllers.MenuController
	store          *store.Store
}

func (this *SnakeGame) Init() tea.Cmd {
	this.store = store.NewStore()

	this.store.Add(consts.HEIGHT, 100)
	this.store.Add(consts.WIDTH, 15)
	this.store.Add(consts.DIFFICULTY, diff_levels.NewEasyLevel())

	firstPage := menu_elements.NewFirstPage(this.store)
	selectableElems := firstPage.SelectableElems()
	firstSelectedElem := selectableElems[0]
	firstSelectedElem.Select()

	this.menuController = controllers.NewMenuController(firstPage)
	this.gameController = controllers.NewGameController()

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
			this.menuController.SelectPrev()
			return this, nil
		case "down":
			this.menuController.SelectNext()
			return this, nil
		case "enter":
			selectedEl := this.menuController.SelectedElem()
			if selectedEl != nil {
				selectedEl.Action(this.store)
				redirectableEl, ok := selectedEl.(game_abstr.IRedirectableElement)
				if ok {
					nextPage := redirectableEl.NextPage(this.store)
					this.menuController.SetPage(nextPage)
				}
			}

			return this, nil
		case "ctrl+c":
		case "esc":
			return this, tea.Quit

		case "w":
			return this, nil
		case "s":
			return this, nil
		case "a":
			return this, nil
		case "d":
			return this, nil
		}

	case tea.MouseMsg:
		mouseMsg := msg.(tea.MouseMsg)
		if mouseMsg.Button == tea.MouseButtonLeft && mouseMsg.Action == tea.MouseActionPress {
		}
		return this, nil

	case tea.WindowSizeMsg:
		width := msg.(tea.WindowSizeMsg).Width
		height := msg.(tea.WindowSizeMsg).Height

		this.store.Add(consts.WIDTH, width)
		this.store.Add(consts.HEIGHT, height)
	}

	return this, nil
}

var _ models.TerminalProgram = (*SnakeGameProgram)(nil)
