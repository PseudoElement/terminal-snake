package menu_elements

import (
	"github.com/charmbracelet/lipgloss"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	consts "github.com/pseudoelement/terminal-snake/src/game/constants"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
)

type DifficultyPage struct {
	store           *store.Store
	selectableElems []game_abstr.ISelectableElement
}

func NewDifficultyPage(store *store.Store) *DifficultyPage {
	selectableElems := []game_abstr.ISelectableElement{
		NewEasyBtn(),
		NewMediumBtn(),
		NewHardBtn(),
	}
	selectableElems[0].Select()

	return &DifficultyPage{selectableElems: selectableElems, store: store}
}

func (this *DifficultyPage) Store() *store.Store {
	return this.store
}

func (this *DifficultyPage) View() string {
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		this.SelectableElemsToViews()...,
	)

	w := this.store.Get(consts.WIDTH).(int)
	h := this.store.Get(consts.HEIGHT).(int)

	flex := lipgloss.Place(
		w, h/2,
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
