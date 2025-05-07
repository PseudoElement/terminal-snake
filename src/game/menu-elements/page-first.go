package menu_elements

import (
	"github.com/charmbracelet/lipgloss"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
)

type FirstPage struct {
	selectableElems []game_abstr.ISelectableElement
}

func NewFirstPage() *FirstPage {
	selectableElems := []game_abstr.ISelectableElement{
		NewPlayBtn(),
		NewQuitBtn(),
	}
	selectableElems[0].Select()

	return &FirstPage{selectableElems: selectableElems}
}

func (this *FirstPage) View() string {
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		this.SelectableElemsToViews()...,
	)

	flex := lipgloss.Place(
		110, 14,
		lipgloss.Center, lipgloss.Bottom,
		content,
	)

	return flex
}

func (this *FirstPage) SelectableElemsToViews() []string {
	var views = make([]string, 0, len(this.selectableElems))
	for _, el := range this.selectableElems {
		views = append(views, el.View())
	}

	return views
}

func (this *FirstPage) SelectableElems() []game_abstr.ISelectableElement {
	return this.selectableElems
}

var _ game_abstr.IPage = (*FirstPage)(nil)
