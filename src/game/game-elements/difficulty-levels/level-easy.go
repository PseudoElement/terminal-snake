package diff_levels

import (
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type EasyLevel struct {
}

func NewEasyLevel() *EasyLevel {
	return &EasyLevel{}
}

func (this *EasyLevel) SnakeSpeedMs() int64 {
	return 170
}

func (this *EasyLevel) Name() string {
	return consts.EASY
}

func (this *EasyLevel) IsSnakeDied(scene game_abstr.IGameScene) bool {
	snakeBody := scene.Snake().Body()
	snakeHead := snakeBody.Head().Val

	return snakeHead.Coords().X < 0 ||
		snakeHead.Coords().Y < 0 ||
		snakeHead.Coords().X >= scene.SceneSize().Width ||
		snakeHead.Coords().Y >= scene.SceneSize().Height
}

var _ game_abstr.IDiffLevel = (*EasyLevel)(nil)
