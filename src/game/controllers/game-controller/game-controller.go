package game_controller

import (
	"time"

	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	menu_elements "github.com/pseudoelement/terminal-snake/src/game/menu-elements"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type GameController struct {
	store           *store.Store
	difficultyLevel game_abstr.IDiffLevel
	gameScene       game_abstr.IGameScene
	showDeathScreen bool
}

func NewGameController(store *store.Store) *GameController {
	gameController := &GameController{
		store:           store,
		difficultyLevel: store.Get(consts.DIFFICULTY).(game_abstr.IDiffLevel),
		showDeathScreen: false,
	}

	return gameController
}

func (this *GameController) SetShowDeathScreen(show bool) {
	this.showDeathScreen = show
}

func (this *GameController) ShowDeathScreen() bool {
	return this.showDeathScreen
}

func (this *GameController) SetGameScene(gameScene game_abstr.IGameScene) {
	this.gameScene = gameScene
}

func (this *GameController) GameScene() game_abstr.IGameScene {
	return this.gameScene
}

func (this *GameController) NextPage() game_abstr.IPage {
	return menu_elements.NewAfterDeathPage(this.store, func() {
		this.SetShowDeathScreen(false)
	})
}

// fix updates view only on button click
func (this *GameController) RunGameLoop() {
	go func() {
		for {
			if this.gameScene.Snake().IsDead() {
				this.SetShowDeathScreen(true)
				return
			}

			delay := time.Duration(this.difficultyLevel.LoopDelayMs())
			time.Sleep(delay * time.Millisecond)

			moveDir := this.store.Get(consts.MOVE_DIRECTION).(game_abstr.MoveDirection)
			snake := this.gameScene.Snake()
			snake.Move(moveDir)
		}
	}()
}
