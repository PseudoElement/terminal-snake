package cell

import (
	"github.com/charmbracelet/lipgloss"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
)

type Cell struct {
	*game_abstr.BaseViewElement
	coords game_abstr.CellCoords
}

func NewCell(cellStyle lipgloss.Style, coords game_abstr.CellCoords) *Cell {
	cell := &Cell{
		BaseViewElement: game_abstr.NewBaseViewElement(cellStyle),
		coords:          coords,
	}
	return cell
}

func (this *Cell) View() string {
	return this.TeaElement().Render()
}

func (this *Cell) Coords() game_abstr.CellCoords {
	return this.coords
}

func (this *Cell) SetCoords(coords game_abstr.CellCoords) {
	this.coords = coords
}

var _ game_abstr.ICell = (*Cell)(nil)
