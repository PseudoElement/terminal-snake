package menu_elements

import (
	"github.com/charmbracelet/lipgloss"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type FirstPage struct {
	store           *store.Store
	selectableElems []game_abstr.ISelectableElement
}

func NewFirstPage(store *store.Store) *FirstPage {
	selectableElems := []game_abstr.ISelectableElement{
		NewPlayBtn(),
		NewDifficultyBtn(),
		NewQuitBtn(),
	}
	selectableElems[0].Select()

	return &FirstPage{selectableElems: selectableElems, store: store}
}

func (this *FirstPage) Store() *store.Store {
	return this.store
}

func (this *FirstPage) View() string {
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
