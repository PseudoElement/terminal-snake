package game_abstr

type CellCoords struct {
	X, Y          int
	Width, Height int
}

type ICell interface {
	IViewElement

	View() string
}

type IGameScene interface {
	GameZone() [][]ICell

	IsSnakeOutScene() bool

	View() string

	Width() int

	Height() int
}
