package menu_elements

import (
	"github.com/charmbracelet/lipgloss"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
)

type DifficultyPage struct {
	selectableElems []game_abstr.ISelectableElement
}

func NewDifficultyPage() *DifficultyPage {
	selectableElems := []game_abstr.ISelectableElement{
		NewEasyBtn(),
		NewHardBtn(),
	}
	selectableElems[0].Select()

	return &DifficultyPage{selectableElems: selectableElems}
}

func (this *DifficultyPage) View() string {
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

func (this *DifficultyPage) SelectableElemsToViews() []string {
	var views = make([]string, 0, len(this.selectableElems))
	for _, el := range this.selectableElems {
		views = append(views, el.View())
	}

	return views
}

func (this *DifficultyPage) SelectableElems() []game_abstr.ISelectableElement {
	return this.selectableElems
}

var _ game_abstr.IPage = (*DifficultyPage)(nil)
