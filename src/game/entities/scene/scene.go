package scene

import (
	"github.com/charmbracelet/lipgloss"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/entities/cell"
	"github.com/pseudoelement/terminal-snake/src/game/entities/snake"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type GameScene struct {
	top, bottom, left, right int
	store                    *store.Store
	snake                    *snake.Snake
	gameZone                 [][]game_abstr.ICell
}

func NewGameScene(store *store.Store) *GameScene {
	bottom := store.Get(consts.HEIGHT).(int)
	right := store.Get(consts.WIDTH).(int)

	scene := &GameScene{
		top:    0,
		bottom: bottom,
		left:   0,
		right:  right,
	}
	scene.createGameZone()
	scene.snake = snake.NewSnake(2, scene, store)

	return scene
}

func (this *GameScene) Width() int {
	return this.right
}

func (this *GameScene) Height() int {
	return this.bottom
}

func (this *GameScene) IsSnakeOutScene() bool {
	snakeBody := this.snake.Body()
	snakeHead := snakeBody.Head().Val

	return snakeHead.Coords().X < this.top ||
		snakeHead.Coords().Y < this.left ||
		snakeHead.Coords().X > this.right ||
		snakeHead.Coords().Y > this.bottom
}

func (this *GameScene) View() string {
	rowsJoins := make([]string, 0, len(this.gameZone)-1)
	for _, row := range this.gameZone {
		cellsToViews := make([]string, 0, len(this.gameZone[0])-1)
		for _, cell := range row {
			cellsToViews = append(cellsToViews, cell.View())
		}

		rowsJoin := lipgloss.JoinHorizontal(
			lipgloss.Center,
			cellsToViews...,
		)
		rowsJoins = append(rowsJoins, rowsJoin)
	}

	verticalJoin := lipgloss.JoinVertical(
		lipgloss.Center,
		rowsJoins...,
	)

	view := lipgloss.Place(
		this.Width()-2, this.Height()-5,
		lipgloss.Center, lipgloss.Center,
		verticalJoin,
	)

	return view
}

func (this *GameScene) GameZone() [][]game_abstr.ICell {
	return this.gameZone
}

func (this *GameScene) createGameZone() {
	rowsCount := this.bottom - 5
	columnsCount := this.right - 5
	rows := make([][]game_abstr.ICell, 0, rowsCount)

	for i := 0; i < rowsCount; i++ {
		row := make([]game_abstr.ICell, 0, columnsCount)
		for j := 0; j < columnsCount; j++ {
			cell := cell.NewCell(cell.GrayCell, cell.CellCoords{
				X: j,
				Y: i,
			})
			row = append(row, cell)
		}

		rows = append(rows, row)
	}

	this.gameZone = rows
}

var _ game_abstr.IGameScene = (*GameScene)(nil)
