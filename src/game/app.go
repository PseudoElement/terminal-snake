package game

import (
	tea "github.com/charmbracelet/bubbletea"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	game_controller "github.com/pseudoelement/terminal-snake/src/game/controllers/game-controller"
	menu_controller "github.com/pseudoelement/terminal-snake/src/game/controllers/menu-controller"
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
	store := store.NewStore()

	store.Set(consts.HEIGHT, 100)
	store.Set(consts.WIDTH, 15)
	store.Set(consts.DIFFICULTY, diff_levels.NewMediumLevel())
	store.Set(consts.MOVE_DIRECTION, consts.RIGHT)
	store.Set(consts.SCORE, 0)

	snakeGame := &SnakeGame{store: store}
	p := tea.NewProgram(
		snakeGame,
		tea.WithAltScreen(),
		// tea.WithMouseCellMotion(),
		tea.WithFPS(30),
	)

	store.Set(consts.PROGRAM, p)

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
	gameController *game_controller.GameController
	menuController *menu_controller.MenuController
	store          *store.Store
}

func (this *SnakeGame) Init() tea.Cmd {
	firstPage := menu_elements.NewFirstPage(this.store)
	selectableElems := firstPage.SelectableElems()
	firstSelectedElem := selectableElems[0]
	firstSelectedElem.Select()

	this.menuController = menu_controller.NewMenuController(firstPage)
	this.gameController = game_controller.NewGameController(this.store)

	return nil
}

// VIEW
func (this *SnakeGame) View() string {
	return this.menuController.Page().View()
}

// UPDATE
func (this *SnakeGame) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case game_abstr.UpdateTrigger:
		return this, nil

	case game_abstr.ShowDeathScreenTrigger:
		afterDeathPage := menu_elements.NewAfterDeathPage(this.store)
		this.menuController.SetPage(afterDeathPage)
		return this, nil

	case game_abstr.RunGameTrigger:
		gamePage, ok := this.menuController.Page().(game_abstr.IGamePage)
		if ok {
			this.gameController.SetGamePage(gamePage)
			this.gameController.RunGame()
		}
		return this, nil

	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {

		case "up":
			if this.menuController.Page().HasSelectableElems() {
				this.menuController.SelectPrev()
			}
			return this, nil
		case "down":
			if this.menuController.Page().HasSelectableElems() {
				this.menuController.SelectNext()
			}
			return this, nil
		case "enter":
			selectedEl := this.menuController.SelectedElem()
			if selectedEl != nil {
				selectedEl.Action(this.store)
				redirectableEl, ok := selectedEl.(game_abstr.IRedirectableElement)
				if ok {
					nextPage := redirectableEl.NextPage(this.store)
					this.menuController.SetPage(nextPage)
					nextPage.OnInit()
				}
			}
			return this, nil
		case "esc":
			firstPage := menu_elements.NewFirstPage(this.store)
			this.menuController.SetPage(firstPage)
			this.gameController.StopGame()
			this.gameController.ResetScore()
			this.store.Set(consts.MOVE_DIRECTION, consts.RIGHT)

			return this, nil
		case "ctrl+c":
			return this, tea.Quit
			// @FIX direction after death
		case "w":
			this.store.Set(consts.MOVE_DIRECTION, consts.UP)
			return this, nil
		case "s":
			this.store.Set(consts.MOVE_DIRECTION, consts.DOWN)
			return this, nil
		case "a":
			this.store.Set(consts.MOVE_DIRECTION, consts.LEFT)
			return this, nil
		case "d":
			this.store.Set(consts.MOVE_DIRECTION, consts.RIGHT)
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

		this.store.Set(consts.WIDTH, width)
		this.store.Set(consts.HEIGHT, height)
	}

	return this, nil
}

var _ models.TerminalProgram = (*SnakeGameProgram)(nil)
