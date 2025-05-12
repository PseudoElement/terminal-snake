package menu_elements

import (
	"github.com/charmbracelet/lipgloss"
	game_abstr "github.com/pseudoelement/terminal-snake/src/game/abstracts"
	"github.com/pseudoelement/terminal-snake/src/game/services/store"
	consts "github.com/pseudoelement/terminal-snake/src/shared/constants"
)

type FirstPage struct {
	*game_abstr.Page
	store *store.Store
}

func NewFirstPage(store *store.Store) *FirstPage {
	selectableElems := []game_abstr.ISelectableElement{
		NewPlayBtn(),
		NewDifficultyBtn(),
		NewQuitBtn(),
	}
	selectableElems[0].Select()

	return &FirstPage{
		Page:  game_abstr.NewPage(store, selectableElems),
		store: store,
	}
}

func (this *FirstPage) View() string {
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

func (this *FirstPage) OnInit() {}

var _ game_abstr.IPage = (*FirstPage)(nil)
