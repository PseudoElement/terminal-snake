package menu_elements

import (
	"github.com/charmbracelet/lipgloss"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type GamePage struct {
	store           *store.Store
	selectableElems []game_abstr.ISelectableElement
}

func NewGamePage(store *store.Store) *GamePage {
	selectableElems := []game_abstr.ISelectableElement{
		NewQuitBtn(),
	}
	selectableElems[0].Select()

	return &GamePage{selectableElems: selectableElems, store: store}
}

func (this *GamePage) Store() *store.Store {
	return this.store
}

func (this *GamePage) View() string {
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		this.SelectableElemsToViews()...,
	)

	w := this.store.Get(consts.WIDTH).(int)
	h := this.store.Get(consts.HEIGHT).(int)

	flex := lipgloss.Place(
		w, h,
		lipgloss.Center, lipgloss.Bottom,
		content,
	)

	return flex
}

func (this *GamePage) SelectableElemsToViews() []string {
	var views = make([]string, 0, len(this.selectableElems))
	for _, el := range this.selectableElems {
		views = append(views, el.View())
	}

	return views
}

func (this *GamePage) SelectableElems() []game_abstr.ISelectableElement {
	return this.selectableElems
}

var _ game_abstr.IPage = (*GamePage)(nil)
