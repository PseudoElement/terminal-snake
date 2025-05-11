package snake

import consts "github.com/pseudoelement/terminal-snake/src/shared/constants"

type MoveDir struct {
	X int
	Y int
}

var directions = map[string]MoveDir{
	consts.UP:    MoveDir{X: 0, Y: -1},
	consts.DOWN:  MoveDir{X: 0, Y: 1},
	consts.LEFT:  MoveDir{X: -1, Y: 0},
	consts.RIGHT: MoveDir{X: 1, Y: 0},
}
