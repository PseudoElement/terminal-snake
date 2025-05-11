package snake

import (
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/entities/cell"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
	data_structs "github.com/pseudoelement/terminal-snake/src/shared/data-structures"
)

type Snake struct {
	body   data_structs.List[*cell.Cell]
	length int
	store  *store.Store
	dead   bool
	scene  game_abstr.IGameScene
}

func NewSnake(startLen int, scene game_abstr.IGameScene, store *store.Store) *Snake {
	if startLen <= 0 {
		startLen = 2
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

func (this *Snake) Body() data_structs.List[*cell.Cell] {
	return this.body
}

func (this *Snake) initBody() {
	body := data_structs.NewList[*cell.Cell](nil)

	terminalWidth := this.store.Get(consts.WIDTH).(int)
	terminalHeight := this.store.Get(consts.HEIGHT).(int)
	xCoord := terminalWidth / 2
	yCoord := terminalHeight / 2

	for i := 0; i < this.length; i++ {
		cell := cell.NewCell(cell.GreenCell, cell.CellCoords{
			X: xCoord,
			Y: yCoord,
		})
		body.Push(cell)
	}

	this.body = body
}

func (this *Snake) Move(direction string) {
	moveDir, ok := directions[direction]
	if !ok {
		panic("invalid direction")
	}

	listNode := this.body.Head()
	prevCoords := this.body.Head().Val.Coords()
	nextCoords := cell.CellCoords{
		X: this.body.Head().Val.Coords().X + moveDir.X,
		Y: this.body.Head().Val.Coords().Y + moveDir.Y,
	}
	for listNode.Next != nil {
		prevCoords = listNode.Val.Coords()
		listNode.Val.SetCoords(nextCoords)
		nextCoords = prevCoords

		listNode = listNode.Next
	}

	if this.scene.IsSnakeOutScene() {
		this.Die()
	}
}

func (this *Snake) Die() {
	this.dead = true
}

func (this *Snake) Eat(c *cell.Cell, scene game_abstr.IGameScene) {
	tail := this.body.Tail()
	preTail := this.body.PreTail()
	tCoords := tail.Val.Coords()
	ptCoords := preTail.Val.Coords()

	newTailCoord := this.defineNewTailCoord(tCoords, ptCoords)
	c.SetCoords(newTailCoord)

	this.body.Push(c)
}

func (this *Snake) defineNewTailCoord(prevTailCoords, prevPreTailCoords cell.CellCoords) cell.CellCoords {
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

	var checkNearestCellAvailable func(tailCoord cell.CellCoords, moveDir MoveDir) (newTailCoord cell.CellCoords, isAvailable bool)
	checkNearestCellAvailable = func(tailCoord cell.CellCoords, moveDir MoveDir) (newTailCoord cell.CellCoords, isAvailable bool) {
		newCoord := cell.CellCoords{
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

	return cell.CellCoords{
		X: prevTailCoords.X + directions[consts.LEFT].X,
		Y: prevTailCoords.Y + directions[consts.LEFT].Y,
	}
}

/*
	x15 y15 -> x16 y15
	x1 y10 -> x0 y10 = x0 y10
*/
