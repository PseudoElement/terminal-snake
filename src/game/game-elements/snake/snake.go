package snake

import (
	"fmt"

	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/game-elements/cell"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
	data_structs "github.com/pseudoelement/terminal-snake/src/shared/data-structures"
)

type Snake struct {
	body   data_structs.List[game_abstr.ICell]
	length int
	store  *store.Store
	dead   bool
	scene  game_abstr.IGameScene
}

func NewSnake(startLen int, scene game_abstr.IGameScene, store *store.Store) *Snake {
	terminalWidth := store.Get(consts.WIDTH).(int)

	if startLen <= 0 {
		startLen = 2
	}
	if startLen >= terminalWidth/2 {
		err := fmt.Sprintf("NewSnake: max len of snake is %d", terminalWidth/2-1)
		panic(err)
	}

	snake := &Snake{
		length: startLen,
		store:  store,
		dead:   false,
		scene:  scene,
	}
	snake.initBody()

	return snake
}

func (this *Snake) Body() data_structs.List[game_abstr.ICell] {
	return this.body
}

func (this *Snake) Find(value game_abstr.CellCoords) (snakeCell game_abstr.ICell, isFound bool) {
	next := this.body.Head()
	for next != nil {
		if next.Val.Coords().X == value.X && next.Val.Coords().Y == value.Y {
			return next.Val, true
		}
		next = next.Next
	}

	return nil, false
}

func (this *Snake) Move(direction string) {
	moveDir, ok := directions[direction]
	if !ok {
		panic("invalid direction")
	}

	listNode := this.body.Head()
	prevCoords := this.body.Head().Val.Coords()
	nextCoords := game_abstr.CellCoords{
		X: this.body.Head().Val.Coords().X + moveDir.X,
		Y: this.body.Head().Val.Coords().Y + moveDir.Y,
	}
	for listNode != nil {
		prevCoords = listNode.Val.Coords()
		listNode.Val.SetCoords(nextCoords)
		nextCoords = prevCoords

		listNode = listNode.Next
	}
}

func (this *Snake) IsDead() bool {
	return this.dead
}

func (this *Snake) Die() {
	this.dead = true
}

func (this *Snake) Eat(c game_abstr.ICell) {
	tail := this.body.Tail()
	preTail := this.body.PreTail()
	tCoords := tail.Val.Coords()
	ptCoords := preTail.Val.Coords()

	newTailCoord := this.defineNewTailCoord(tCoords, ptCoords)
	newTailCell := cell.NewCell(cell.BlueCell, newTailCoord)

	this.body.Push(newTailCell)
}

func (this *Snake) defineNewTailCoord(prevTailCoords, prevPreTailCoords game_abstr.CellCoords) game_abstr.CellCoords {
	var priorityDirection = MoveDir{X: 1, Y: 0}
	if prevTailCoords.X > prevPreTailCoords.X {
		// to right
		priorityDirection = MoveDir{X: 1, Y: 0}
	} else if prevTailCoords.X < prevPreTailCoords.X {
		// to left
		priorityDirection = MoveDir{X: -1, Y: 0}
	} else if prevTailCoords.Y > prevPreTailCoords.Y {
		// to bottom
		priorityDirection = MoveDir{X: 0, Y: 1}
	} else {
		//to top
		priorityDirection = MoveDir{X: 0, Y: -1}
	}

	var checkNearestCellAvailable func(tailCoord game_abstr.CellCoords, moveDir MoveDir) (newTailCoord game_abstr.CellCoords, isAvailable bool)
	checkNearestCellAvailable = func(tailCoord game_abstr.CellCoords, moveDir MoveDir) (newTailCoord game_abstr.CellCoords, isAvailable bool) {
		newCoord := game_abstr.CellCoords{
			X: tailCoord.X + moveDir.X,
			Y: tailCoord.Y + moveDir.Y,
		}

		yHeight := len(this.scene.GameZone()) - 1
		xWidth := len(this.scene.GameZone()[0]) - 1

		var available bool
		if newCoord.X >= 0 && newCoord.Y >= 0 && newCoord.X <= xWidth && newCoord.Y <= yHeight {
			available = true
		} else {
			available = false
		}

		return newCoord, available
	}

	newTailCoord, available := checkNearestCellAvailable(prevPreTailCoords, priorityDirection)
	if available {
		return newTailCoord
	}

	for _, moveDir := range directions {
		if moveDir == priorityDirection {
			continue
		}
		newTailCoord, available := checkNearestCellAvailable(prevPreTailCoords, priorityDirection)
		if available {
			return newTailCoord
		}
	}

	return game_abstr.CellCoords{
		X: prevTailCoords.X + directions[consts.LEFT].X,
		Y: prevTailCoords.Y + directions[consts.LEFT].Y,
	}
}

func (this *Snake) initBody() {
	xCoord := this.scene.SceneSize().Width / 2
	yCoord := this.scene.SceneSize().Height / 2

	firstCell := &data_structs.ListNode[game_abstr.ICell]{
		Val: cell.NewCell(cell.PinkCell, game_abstr.CellCoords{
			X: xCoord,
			Y: yCoord,
		}),
		Next: nil,
	}
	body := data_structs.NewList(firstCell)

	for i := 1; i < this.length; i++ {
		cell := cell.NewCell(cell.BlueCell, game_abstr.CellCoords{
			X: xCoord - i,
			Y: yCoord,
		})
		body.Push(cell)
	}

	this.body = body
}

var _ game_abstr.ISnake = (*Snake)(nil)
