package scene

import (
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/charmbracelet/lipgloss"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/game-elements/cell"
	"github.com/pseudoelement/terminal-snake/src/game/game-elements/food"
	"github.com/pseudoelement/terminal-snake/src/game/game-elements/snake"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type GameScene struct {
	sceneSize game_abstr.SceneSize
	store     *store.Store
	snake     *snake.Snake
	gameZone  [][]game_abstr.ICell
	food      game_abstr.ICell
}

func NewGameScene(store *store.Store) *GameScene {
	h := store.Get(consts.HEIGHT).(int)
	w := store.Get(consts.WIDTH).(int)
	sceneHeight := h - 5
	sceneWidth := int(math.Floor(float64(w) / 2.1))

	scene := &GameScene{
		store: store,
		sceneSize: game_abstr.SceneSize{
			Width:  sceneWidth,
			Height: sceneHeight,
		},
		food: nil,
	}
	scene.createGameZone()
	scene.SpawnFood()
	scene.snake = snake.NewSnake(2, scene, store)

	return scene
}

func (this *GameScene) Snake() game_abstr.ISnake {
	return this.snake
}

func (this *GameScene) SceneSize() game_abstr.SceneSize {
	return this.sceneSize
}

func (this *GameScene) Food() game_abstr.ICell {
	return this.food
}

func (this *GameScene) DoesSnakeTakeFood() bool {
	body := this.snake.Body()
	headCoords := body.Head().Val.Coords()
	foodCoords := this.food.Coords()

	return headCoords.X == foodCoords.X && headCoords.Y == foodCoords.Y
}

func (this *GameScene) IsSnakeOutScene() bool {
	snakeBody := this.snake.Body()
	snakeHead := snakeBody.Head().Val

	return snakeHead.Coords().X < 0 ||
		snakeHead.Coords().Y < 0 ||
		snakeHead.Coords().X > this.sceneSize.Width ||
		snakeHead.Coords().Y > this.sceneSize.Height
}

func (this *GameScene) SpawnFood() {
	randomCoords := game_abstr.CellCoords{
		X: rand.IntN(this.sceneSize.Width),
		Y: rand.IntN(this.sceneSize.Height),
	}
	this.food = food.NewFood(randomCoords)
}

func (this *GameScene) RemoveFood() {
	this.food = nil
}

func (this *GameScene) View() string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Scene.View(). Error:\n", r)
		}
	}()

	rowsJoins := make([]string, 0, len(this.gameZone)-1)
	for y, row := range this.gameZone {
		cellsToViews := make([]string, 0, len(this.gameZone[0])-1)
		for x, rowCell := range row {
			snakeCell, found := this.snake.Find(game_abstr.CellCoords{X: x, Y: y})
			if found {
				cellsToViews = append(cellsToViews, snakeCell.View())
			} else if this.food.Coords().X == x && this.food.Coords().Y == y {
				cellsToViews = append(cellsToViews, this.food.View())
			} else {
				cellsToViews = append(cellsToViews, rowCell.View())
			}
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
		this.sceneSize.Width-2, this.sceneSize.Height-5,
		lipgloss.Center, lipgloss.Center,
		verticalJoin,
	)

	return view
}

func (this *GameScene) GameZone() [][]game_abstr.ICell {
	return this.gameZone
}

func (this *GameScene) createGameZone() {
	rowsCount := this.sceneSize.Height
	columnsCount := this.sceneSize.Width
	rows := make([][]game_abstr.ICell, 0, rowsCount)

	for i := 0; i < rowsCount; i++ {
		row := make([]game_abstr.ICell, 0, columnsCount)
		for j := 0; j < columnsCount; j++ {
			if (i%2 == 0 && j%2 == 0) || (i%2 != 0 && j%2 != 0) {
				cell := cell.NewCell(cell.DarkGrayCell, game_abstr.CellCoords{
					X: j,
					Y: i,
				})
				row = append(row, cell)
			} else {
				cell := cell.NewCell(cell.GrayCell, game_abstr.CellCoords{
					X: j,
					Y: i,
				})
				row = append(row, cell)
			}
		}

		rows = append(rows, row)
	}

	this.gameZone = rows
}

var _ game_abstr.IGameScene = (*GameScene)(nil)
