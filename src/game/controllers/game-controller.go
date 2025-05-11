package controllers

type GameController struct {
	isPlaying bool
}

func NewGameController() *GameController {
	return &GameController{isPlaying: false}
}

func (this *GameController) IsPlaying() bool {
	return this.isPlaying
}
