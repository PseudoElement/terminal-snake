package food

import (
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/game-elements/cell"
)

type Food struct {
	*cell.Cell
}

func NewFood(coords game_abstr.CellCoords) *Food {
	return &Food{
		Cell: cell.NewCell(cell.RedCell, coords),
	}
}
