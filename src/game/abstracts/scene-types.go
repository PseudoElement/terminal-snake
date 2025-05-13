package game_abstr

import (
	data_structs "github.com/pseudoelement/terminal-snake/src/shared/data-structures"
)

type MoveDirection = string

type CellCoords struct {
	X, Y int
}

type SceneSize struct {
	Width  int
	Height int
}

type ICell interface {
	IViewElement

	Coords() CellCoords

	SetCoords(cellCoord CellCoords)
}

type IGameScene interface {
	GameZone() [][]ICell

	DoesSnakeTakeFood() bool

	View() string

	SceneSize() SceneSize

	Snake() ISnake

	SpawnFood()

	RemoveFood()

	Food() ICell
}

type ISnake interface {
	Body() data_structs.List[ICell]

	Find(value CellCoords) (snakeCell ICell, isFound bool)

	Move(direction MoveDirection)

	Die()

	Eat(c ICell)

	IsDead() bool
}
