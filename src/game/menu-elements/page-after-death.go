package menu_elements

import (
	"github.com/charmbracelet/lipgloss"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type AfterDeathPage struct {
	*game_abstr.Page
}

func NewAfterDeathPage(store *store.Store, closePage func()) *AfterDeathPage {
	selectableElems := []game_abstr.ISelectableElement{
		NewBackToMenuBtn(closePage),
	}
	selectableElems[0].Select()

	return &AfterDeathPage{
		Page: game_abstr.NewPage(store, selectableElems),
	}
}

func (this *AfterDeathPage) View() string {
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		this.SelectableElemsToViews()...,
	)

	w := this.Store().Get(consts.WIDTH).(int)
	h := this.Store().Get(consts.HEIGHT).(int)

	flex := lipgloss.Place(
		w, h/2,
		lipgloss.Center, lipgloss.Bottom,
		content,
	)

	return flex
}

var _ game_abstr.IPage = (*AfterDeathPage)(nil)
