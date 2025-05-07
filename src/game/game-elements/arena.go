package game_elements

type GameArena struct{}

func NewGameArena() *GameArena {
	return &GameArena{}
}

// var _ game_abstr.IViewElement = (*GameArena)(nil)
