package game_controller

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	menu_elements "github.com/pseudoelement/terminal-snake/src/game/menu-elements"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type GameController struct {
	store      *store.Store
	teaProgram *tea.Program
	stop       bool
	gamePage   game_abstr.IGamePage
	gameScene  game_abstr.IGameScene
}

func NewGameController(store *store.Store) *GameController {
	gameController := &GameController{
		store:      store,
		teaProgram: store.Get(consts.PROGRAM).(*tea.Program),
		stop:       false,
		gamePage:   nil,
	}

	return gameController
}

func (this *GameController) SetGamePage(gamePage game_abstr.IGamePage) {
	this.gamePage = gamePage
	this.gameScene = gamePage.GameScene()
}

func (this *GameController) NextPage() game_abstr.IPage {
	return menu_elements.NewAfterDeathPage(this.store)
}

func (this *GameController) StopGame() {
	this.stop = true
}

func (this *GameController) RunGame() {
	this.stop = false

	go func() {
		for !this.gameScene.Snake().IsDead() && !this.stop {
			diffLevel := this.store.Get(consts.DIFFICULTY).(game_abstr.IDiffLevel)
			delay := time.Duration(diffLevel.LoopDelayMs())
			time.Sleep(delay * time.Millisecond)

			moveDir := this.store.Get(consts.MOVE_DIRECTION).(game_abstr.MoveDirection)
			snake := this.gameScene.Snake()
			snake.Move(moveDir)

			if diffLevel.IsSnakeDied(this.gameScene) {
				snake.Die()
				break
			}
			if this.gameScene.DoesSnakeTakeFood() {
				snake.Eat(this.gameScene.Food())
				this.IncrementScore()
				this.gameScene.RemoveFood()
				this.gameScene.SpawnFood()
			}

			this.teaProgram.Send(game_abstr.UpdateTrigger{})
		}

		if this.gameScene.Snake().IsDead() {
			this.teaProgram.Send(game_abstr.ShowDeathScreenTrigger{})
		}
		this.StopGame()
	}()
}

func (this *GameController) IncrementScore() {
	currScore := this.store.Get(consts.SCORE).(int)
	currScore++
	this.store.Set(consts.SCORE, currScore)
}

func (this *GameController) ResetScore() {
	this.store.Set(consts.SCORE, 0)
}
