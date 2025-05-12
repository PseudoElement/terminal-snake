package diff_levels

import (
	"fmt"

	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
)

type HardLevel struct {
}

func NewHardLevel() *HardLevel {
	return &HardLevel{}
}

func (this *HardLevel) LoopDelayMs() int64 {
	return 50
}

func (this *HardLevel) IsSnakeDied(scene game_abstr.IGameScene) bool {
	snakeBody := scene.Snake().Body()
	snakeHead := snakeBody.Head().Val

	visited := make(map[string]bool, snakeBody.Size())
	next := snakeBody.Head()
	for next.Next != nil {
		key := fmt.Sprintf("%v_%v", next.Val.Coords().X, next.Val.Coords().Y)
		_, crossed := visited[key]
		if crossed {
			return true
		}

		visited[key] = true
		next = next.Next
	}

	return snakeHead.Coords().X < 0 ||
		snakeHead.Coords().Y < 0 ||
		snakeHead.Coords().X >= scene.SceneSize().Width ||
		snakeHead.Coords().Y >= scene.SceneSize().Height
}

var _ game_abstr.IDiffLevel = (*HardLevel)(nil)
