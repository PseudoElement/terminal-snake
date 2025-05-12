package menu_elements

import (
	"github.com/charmbracelet/lipgloss"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type RulesPage struct {
	*game_abstr.Page
	store *store.Store
}

func NewRulesPage(store *store.Store) *RulesPage {
	selectableElems := []game_abstr.ISelectableElement{
		NewBackToMenuBtn(),
	}
	selectableElems[0].Select()

	return &RulesPage{
		Page:  game_abstr.NewPage(store, selectableElems),
		store: store,
	}
}

func (this *RulesPage) View() string {
	content := append([]string{NewTextRules().View()}, this.SelectableElemsToViews()...)

	joinVertical := lipgloss.JoinVertical(
		lipgloss.Center,
		content...,
	)

	w := this.Store().Get(consts.WIDTH).(int)
	h := this.Store().Get(consts.HEIGHT).(int)

	flex := lipgloss.Place(
		w, h/2,
		lipgloss.Center, lipgloss.Bottom,
		joinVertical,
	)

	return flex
}

func (this *RulesPage) OnInit() {}

var _ game_abstr.IPage = (*RulesPage)(nil)
