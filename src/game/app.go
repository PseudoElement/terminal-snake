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
	p := tea.NewProgram(
		&SnakeGame{},
		tea.WithAltScreen(),
		// tea.WithMouseCellMotion(),
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
	gameController *game_controller.GameController
	menuController *menu_controller.MenuController
	store          *store.Store
}

func (this *SnakeGame) Init() tea.Cmd {
	this.store = store.NewStore()

	this.store.Add(consts.HEIGHT, 100)
	this.store.Add(consts.WIDTH, 15)
	this.store.Add(consts.DIFFICULTY, diff_levels.NewEasyLevel())
	this.store.Add(consts.MOVE_DIRECTION, consts.RIGHT)

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
	if this.gameController != nil && this.gameController.ShowDeathScreen() {
		adp := menu_elements.NewAfterDeathPage(this.store, func() {
			this.gameController.SetShowDeathScreen(false)
		})
		this.menuController.SetPage(adp)
		return adp.View()
	}
	return this.menuController.Page().View()
}

// UPDATE
func (this *SnakeGame) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
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
					gamePage, ok := nextPage.(game_abstr.IGamePage)
					if ok {
						this.gameController.SetGameScene(gamePage.GameScene())
						this.gameController.RunGameLoop()
					}
				}
			}
			return this, nil
		case "esc":
			firstPage := menu_elements.NewFirstPage(this.store)
			this.menuController.SetPage(firstPage)
			return this, nil
		case "ctrl+c":
			return this, tea.Quit

		case "w":
			this.store.Add(consts.MOVE_DIRECTION, consts.UP)
			return this, nil
		case "s":
			this.store.Add(consts.MOVE_DIRECTION, consts.DOWN)
			return this, nil
		case "a":
			this.store.Add(consts.MOVE_DIRECTION, consts.LEFT)
			return this, nil
		case "d":
			this.store.Add(consts.MOVE_DIRECTION, consts.RIGHT)
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
