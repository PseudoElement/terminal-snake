package cell

import (
	"github.com/charmbracelet/lipgloss"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
)

type CellCoords struct {
	X, Y int
}

type Cell struct {
	game_abstr.BaseViewElement
	coords CellCoords
}

func NewCell(cellStyle lipgloss.Style, coords CellCoords) *Cell {
	cell := &Cell{
		BaseViewElement: game_abstr.NewBaseViewElement(cellStyle),
		coords:          coords,
	}
	return cell
}

func (this *Cell) View() string {
	return this.TeaElement().Render()
}

func (this *Cell) Coords() CellCoords {
	return this.coords
}

func (this *Cell) SetCoords(coords CellCoords) {
	this.coords = coords
}

// var _ game_abstr.IViewElement = (*Cell)(nil)
